package bootstrap

import (
	authService "mytodo/apps/api/internal/auth/domain/service"
	authPersistence "mytodo/apps/api/internal/auth/infrastructure/persistence"
	authHttp "mytodo/apps/api/internal/auth/interfaces/http"
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
	userRepo := authPersistence.NewUserRepository(db.DB)
	sessionRepo := authPersistence.NewSessionRepository(db.DB)

	// Initialize auth service
	authSvc := authService.NewAuthService(userRepo, sessionRepo, jwtService, passwordService)

	// Initialize controllers
	authController := authHttp.NewAuthController(authSvc)

	return &Container{
		DB:              db,
		Redis:           redisClient,
		Log:             logger,
		Config:          cfg,
		JWTService:      jwtService,
		PasswordService: passwordService,
		AuthService:     authSvc,
		AuthController:  authController,
	}, nil
}
