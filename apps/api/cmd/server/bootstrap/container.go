package bootstrap

import (
	authService "mytodo/apps/api/internal/auth/domain/service"
	authPersistence "mytodo/apps/api/internal/auth/infrastructure/persistence"
	authGrpc "mytodo/apps/api/internal/auth/interfaces/grpc"
	authHttp "mytodo/apps/api/internal/auth/interfaces/http"

	userHandlers "mytodo/apps/api/internal/users/application/handlers"
	userPersistence "mytodo/apps/api/internal/users/infrastructure/persistence"
	userHttp "mytodo/apps/api/internal/users/interfaces/http"

	"mytodo/apps/api/pkg/cache/redis"
	"mytodo/apps/api/pkg/database/postgres"
	"mytodo/apps/api/pkg/security"
)

type Container struct {
	DB              *postgres.DB
	Redis           *redis.Client
	Log             Logger
	Config          *Config
	JWTService      *security.JWTService
	PasswordService *security.PasswordService
	AuthService     *authService.AuthService
	AuthController  *authHttp.AuthController
	AuthGrpcServer  *authGrpc.AuthServer
	UserController  *userHttp.UserController
}

func NewContainer(logger Logger, cfg *Config) (*Container, error) {

	db, err := postgres.NewPostgresConnection()
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.NewRedisClient()
	if err != nil {
		return nil, err
	}

	// Initialize security services
	jwtService := security.NewJWTService(cfg.JWTSecret, cfg.JWTExpiry, 24*7) // 7 days for refresh
	passwordService := security.NewPasswordService()

	// Initialize repositories
	authUserRepo := authPersistence.NewUserRepository(db.DB)
	sessionRepo := authPersistence.NewSessionRepository(db.DB)
	userProfileRepo := userPersistence.NewUserRepository(db.DB)

	// Initialize auth service
	authSvc := authService.NewAuthService(authUserRepo, sessionRepo, jwtService, passwordService)

	// Initialize user handlers and controller
	userQueryHandler := userHandlers.NewUserQueryHandler(userProfileRepo)
	userCommandHandler := userHandlers.NewUserCommandHandler(userProfileRepo)
	userController := userHttp.NewUserController(userQueryHandler, userCommandHandler)

	registrationSvc := authService.NewUserRegistrationService(
		authSvc,
		userCommandHandler,
		authUserRepo,
		sessionRepo,
		jwtService,
		passwordService,
	)

	// Initialize controllers
	authController := authHttp.NewAuthController(authSvc, registrationSvc)

	// Initialize gRPC servers
	authGrpcServer := authGrpc.NewAuthServer(authSvc)

	return &Container{
		DB:              db,
		Redis:           redisClient,
		Log:             logger,
		Config:          cfg,
		JWTService:      jwtService,
		PasswordService: passwordService,
		AuthService:     authSvc,
		AuthController:  authController,
		AuthGrpcServer:  authGrpcServer,
		UserController:  userController,
	}, nil
}
