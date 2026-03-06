package bootstrap

import (
	"mytodo/apps/api/pkg/database"
)

type Container struct {
	Config *Config
	Logger Logger
	DB     *database.DB
	// Cache  *cache.RedisClient

	// // Repositories
	// UserRepo    users.UserRepository
	// ProjectRepo projects.ProjectRepository
	// IssueRepo   issues.IssueRepository

	// // Services
	// AuthService auth.AuthService
	// UserService users.UserService

	// // Handlers
	// AuthHandler    *auth.AuthHandler
	// UserHandler    *users.UserHandler
	// ProjectHandler *projects.ProjectHandler
	// IssueHandler   *issues.IssueHandler
}

func NewContainer(config *Config, logger Logger) *Container {
	// Initialize database connection
	db := database.NewDb(database.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DBName:   config.DBName,
	})

	// Initialize cache
	// cache := cache.NewRedisClient(cache.Config{
	// 	Host: config.RedisHost,
	// 	Port: config.RedisPort,
	// 	Pass: config.RedisPass,
	// })

	// // Initialize repositories
	// userRepo := users.NewUserRepository(db)
	// projectRepo := projects.NewProjectRepository(db)
	// issueRepo := issues.NewIssueRepository(db)

	// // Initialize services
	// authService := auth.NewAuthService(userRepo, cache, config.JWTSecret, config.JWTExpiry)
	// userService := users.NewUserService(userRepo)

	// // Initialize handlers
	// authHandler := auth.NewAuthHandler(authService, logger)
	// userHandler := users.NewUserHandler(userService, logger)
	// projectHandler := projects.NewProjectHandler(projectRepo, logger)
	// issueHandler := issues.NewIssueHandler(issueRepo, projectRepo, logger)

	return &Container{
		Config: config,
		Logger: logger,
		DB:     db,
		// Cache:          cache,
		// UserRepo:       userRepo,
		// ProjectRepo:    projectRepo,
		// IssueRepo:      issueRepo,
		// AuthService:    authService,
		// UserService:    userService,
		// AuthHandler:    authHandler,
		// UserHandler:    userHandler,
		// ProjectHandler: projectHandler,
		// IssueHandler:   issueHandler,
	}
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
	if c.Cache != nil {
		c.Cache.Close()
	}
}
