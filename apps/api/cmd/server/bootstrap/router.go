package bootstrap

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Package bootstrap handles application initialization.
//
// This file sets up HTTP routing and middleware configuration.
//
// In production-grade applications, router setup typically includes:
// - Route grouping and versioning (/api/v1, /api/v2)
// - Middleware chaining (logging, auth, CORS, rate limiting)
// - Route-level middleware (specific auth for endpoints)
// - Static file serving
// - WebSocket routes
// - OpenAPI/Swagger documentation endpoint
// - Health check and metrics endpoints
//
// Example structure:
//   func SetupRouter(container *Container) *mux.Router {
//       router := mux.NewRouter()
//
//       // Global middleware
//       router.Use(middleware.RequestID)
//       router.Use(middleware.Logger)
//       router.Use(middleware.Recovery)
//       router.Use(middleware.CORS)
//
//       // Health endpoints (no auth)
//       router.HandleFunc("/health", healthHandler).Methods("GET")
//       router.HandleFunc("/ready", readinessHandler).Methods("GET")
//
//       // API v1 routes
//       apiV1 := router.PathPrefix("/api/v1").Subrouter()
//       apiV1.Use(middleware.Auth)
//
//       // Register module routes
//       registerAuthRoutes(apiV1, container)
//       registerProjectRoutes(apiV1, container)
//       registerIssueRoutes(apiV1, container)
//
//       return router
//   }
//
// Example middleware chain:
//   Request -> RequestID -> Logger -> Recovery -> CORS -> Auth -> Handler

type Router struct {
	mux    *http.ServeMux
	app    *App
	logger Logger
}

func NewRouter(app *App, logger Logger) *Router {
	router := &Router{
		mux:    http.NewServeMux(),
		app:    app,
		logger: logger,
	}

	router.setupRoutes()
	return router
}

func (r *Router) setupRoutes() {
	// Health check
	r.mux.HandleFunc("GET /health", r.healthCheck)

	// API v1 routes
	// Auth routes
	r.mux.HandleFunc("POST /api/v1/auth/register", r.handleAuthRegister)
	r.mux.HandleFunc("POST /api/v1/auth/login", r.handleAuthLogin)

	// User routes (protected)
	r.mux.HandleFunc("GET /api/v1/users/profile", r.authMiddleware(r.handleGetProfile))
	r.mux.HandleFunc("PUT /api/v1/users/profile", r.authMiddleware(r.handleUpdateProfile))

	// Project routes (protected)
	r.mux.HandleFunc("POST /api/v1/projects", r.authMiddleware(r.handleCreateProject))
	r.mux.HandleFunc("GET /api/v1/projects", r.authMiddleware(r.handleListProjects))
	r.mux.HandleFunc("GET /api/v1/projects/", r.authMiddleware(r.handleGetProject))
	r.mux.HandleFunc("PUT /api/v1/projects/", r.authMiddleware(r.handleUpdateProject))
	r.mux.HandleFunc("DELETE /api/v1/projects/", r.authMiddleware(r.handleDeleteProject))

	// Issue routes (protected)
	r.mux.HandleFunc("POST /api/v1/issues", r.authMiddleware(r.handleCreateIssue))
	r.mux.HandleFunc("GET /api/v1/issues", r.authMiddleware(r.handleListIssues))
	r.mux.HandleFunc("GET /api/v1/issues/", r.authMiddleware(r.handleGetIssue))
	r.mux.HandleFunc("PUT /api/v1/issues/", r.authMiddleware(r.handleUpdateIssue))
	r.mux.HandleFunc("DELETE /api/v1/issues/", r.authMiddleware(r.handleDeleteIssue))
}

func (r *Router) Run(addr string) error {
	r.logger.Info("Router listening on " + addr)
	return http.ListenAndServe(addr, r.mux)
}

func (r *Router) Shutdown(ctx context.Context) error {
	r.logger.Info("Shutting down router...")
	return nil
}

// Middleware
func (r *Router) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		token := extractToken(req.Header.Get("Authorization"))
		if token == "" {
			r.respondWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		// Validate token (simplified - in real app, validate with auth service)
		userID, err := r.app.GetAuthHandler().ValidateToken(token)
		if err != nil {
			r.respondWithError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		// Add user ID to request context
		ctx := req.Context()
		ctx = context.WithValue(ctx, "userID", userID)
		next.ServeHTTP(w, req.WithContext(ctx))
	}
}

// Health check
func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
	r.respondWithJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"time":   time.Now().String(),
	})
}

// Auth handlers
func (r *Router) handleAuthRegister(w http.ResponseWriter, req *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := r.parseJSON(req, &input); err != nil {
		r.respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	user, err := r.app.GetAuthHandler().Register(input.Email, input.Password, input.Name)
	if err != nil {
		r.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.respondWithJSON(w, http.StatusCreated, user)
}

func (r *Router) handleAuthLogin(w http.ResponseWriter, req *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := r.parseJSON(req, &input); err != nil {
		r.respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	token, err := r.app.GetAuthHandler().Login(input.Email, input.Password)
	if err != nil {
		r.respondWithError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	r.respondWithJSON(w, http.StatusOK, map[string]string{
		"token": token,
		"type":  "Bearer",
	})
}

// User handlers
func (r *Router) handleGetProfile(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("userID").(int)
	user, err := r.app.GetUserHandler().GetByID(userID)
	if err != nil {
		r.respondWithError(w, http.StatusNotFound, "user not found")
		return
	}

	r.respondWithJSON(w, http.StatusOK, user)
}

func (r *Router) handleUpdateProfile(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("userID").(int)

	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := r.parseJSON(req, &input); err != nil {
		r.respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	user, err := r.app.GetUserHandler().Update(userID, input.Name, input.Email)
	if err != nil {
		r.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.respondWithJSON(w, http.StatusOK, user)
}

// Project handlers
func (r *Router) handleCreateProject(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("userID").(int)

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := r.parseJSON(req, &input); err != nil {
		r.respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	project, err := r.app.GetProjectHandler().Create(userID, input.Name, input.Description)
	if err != nil {
		r.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.respondWithJSON(w, http.StatusCreated, project)
}

func (r *Router) handleListProjects(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("userID").(int)
	projects, err := r.app.GetProjectHandler().ListByUser(userID)
	if err != nil {
		r.respondWithError(w, http.StatusInternalServerError, "failed to list projects")
		return
	}

	r.respondWithJSON(w, http.StatusOK, projects)
}

func (r *Router) handleGetProject(w http.ResponseWriter, req *http.Request) {
	id := extractID(req.URL.Path)
	if id == 0 {
		r.respondWithError(w, http.StatusBadRequest, "invalid project ID")
		return
	}

	project, err := r.app.GetProjectHandler().GetByID(id)
	if err != nil {
		r.respondWithError(w, http.StatusNotFound, "project not found")
		return
	}

	r.respondWithJSON(w, http.StatusOK, project)
}

func (r *Router) handleUpdateProject(w http.ResponseWriter, req *http.Request) {
	id := extractID(req.URL.Path)
	if id == 0 {
		r.respondWithError(w, http.StatusBadRequest, "invalid project ID")
		return
	}

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := r.parseJSON(req, &input); err != nil {
		r.respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	project, err := r.app.GetProjectHandler().Update(id, input.Name, input.Description)
	if err != nil {
		r.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.respondWithJSON(w, http.StatusOK, project)
}

func (r *Router) handleDeleteProject(w http.ResponseWriter, req *http.Request) {
	id := extractID(req.URL.Path)
	if id == 0 {
		r.respondWithError(w, http.StatusBadRequest, "invalid project ID")
		return
	}

	if err := r.app.GetProjectHandler().Delete(id); err != nil {
		r.respondWithError(w, http.StatusNotFound, "project not found")
		return
	}

	r.respondWithJSON(w, http.StatusNoContent, nil)
}

// Issue handlers
func (r *Router) handleCreateIssue(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("userID").(int)

	var input struct {
		ProjectID   int    `json:"project_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := r.parseJSON(req, &input); err != nil {
		r.respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	issue, err := r.app.GetIssueHandler().Create(userID, input.ProjectID, input.Title, input.Description)
	if err != nil {
		r.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.respondWithJSON(w, http.StatusCreated, issue)
}

func (r *Router) handleListIssues(w http.ResponseWriter, req *http.Request) {
	projectID := req.URL.Query().Get("project_id")
	var issues interface{}
	var err error

	if projectID != "" {
		id := parseInt(projectID)
		issues, err = r.app.GetIssueHandler().ListByProject(id)
	} else {
		userID := req.Context().Value("userID").(int)
		issues, err = r.app.GetIssueHandler().ListByUser(userID)
	}

	if err != nil {
		r.respondWithError(w, http.StatusInternalServerError, "failed to list issues")
		return
	}

	r.respondWithJSON(w, http.StatusOK, issues)
}

func (r *Router) handleGetIssue(w http.ResponseWriter, req *http.Request) {
	id := extractID(req.URL.Path)
	if id == 0 {
		r.respondWithError(w, http.StatusBadRequest, "invalid issue ID")
		return
	}

	issue, err := r.app.GetIssueHandler().GetByID(id)
	if err != nil {
		r.respondWithError(w, http.StatusNotFound, "issue not found")
		return
	}

	r.respondWithJSON(w, http.StatusOK, issue)
}

func (r *Router) handleUpdateIssue(w http.ResponseWriter, req *http.Request) {
	id := extractID(req.URL.Path)
	if id == 0 {
		r.respondWithError(w, http.StatusBadRequest, "invalid issue ID")
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	if err := r.parseJSON(req, &input); err != nil {
		r.respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	issue, err := r.app.GetIssueHandler().Update(id, input.Title, input.Description, input.Status)
	if err != nil {
		r.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.respondWithJSON(w, http.StatusOK, issue)
}

func (r *Router) handleDeleteIssue(w http.ResponseWriter, req *http.Request) {
	id := extractID(req.URL.Path)
	if id == 0 {
		r.respondWithError(w, http.StatusBadRequest, "invalid issue ID")
		return
	}

	if err := r.app.GetIssueHandler().Delete(id); err != nil {
		r.respondWithError(w, http.StatusNotFound, "issue not found")
		return
	}

	r.respondWithJSON(w, http.StatusNoContent, nil)
}

// Helper functions
func (r *Router) parseJSON(req *http.Request, v interface{}) error {
	return json.NewDecoder(req.Body).Decode(v)
}

func (r *Router) respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (r *Router) respondWithError(w http.ResponseWriter, status int, message string) {
	r.respondWithJSON(w, status, map[string]string{"error": message})
}

func extractToken(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}

	return parts[1]
}

func extractID(path string) int {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return 0
	}

	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0
	}

	return id
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
