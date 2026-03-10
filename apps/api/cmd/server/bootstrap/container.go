package bootstrap

import (
	authService "mytodo/apps/api/internal/auth/domain/service"
	authPersistence "mytodo/apps/api/internal/auth/infrastructure/persistence"
	authGrpc "mytodo/apps/api/internal/auth/interfaces/grpc"
	authHttp "mytodo/apps/api/internal/auth/interfaces/http"

	orgHandlers "mytodo/apps/api/internal/organizations/application/handlers"
	orgService "mytodo/apps/api/internal/organizations/domain/service"
	orgPersistence "mytodo/apps/api/internal/organizations/infrastructure/persistence"
	orgHttp "mytodo/apps/api/internal/organizations/interfaces/http"

	userHandlers "mytodo/apps/api/internal/users/application/handlers"
	userService "mytodo/apps/api/internal/users/domain/service"
	userPersistence "mytodo/apps/api/internal/users/infrastructure/persistence"
	userHttp "mytodo/apps/api/internal/users/interfaces/http"

	"mytodo/apps/api/pkg/cache/redis"
	"mytodo/apps/api/pkg/database/postgres"
	"mytodo/apps/api/pkg/security"
)

type Container struct {
	DB               *postgres.DB
	Redis            *redis.Client
	Log              Logger
	Config           *Config
	JWTService       *security.JWTService
	PasswordService  *security.PasswordService
	AuthService      *authService.AuthService
	AuthController   *authHttp.AuthController
	AuthGrpcServer   *authGrpc.AuthServer
	UserController   *userHttp.UserController
	OrgController    *orgHttp.OrganizationController
	MemberController *orgHttp.MemberController
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
	orgRepo := orgPersistence.NewOrganizationRepository(db.DB)
	membershipRepo := orgPersistence.NewMembershipRepository(db.DB)

	profileSvc := userService.NewProfileService(userProfileRepo)
	userSvc := userService.NewUserService(userProfileRepo)

	// Initialize auth service
	authSvc := authService.NewAuthService(authUserRepo, sessionRepo, jwtService, passwordService)

	// Initialize user handlers and controller
	userQueryHandler := userHandlers.NewUserQueryHandler(userSvc)
	userCommandHandler := userHandlers.NewUserCommandHandler(profileSvc)
	userController := userHttp.NewUserController(userQueryHandler, userCommandHandler)

	// Initialize organization service, handlers, and controller
	orgSvc := orgService.NewOrganizationService(orgRepo)
	orgQueryHandler := orgHandlers.NewOrganizationQueryHandler(orgSvc)
	orgCommandHandler := orgHandlers.NewOrganizationCommandHandler(orgSvc)
	orgController := orgHttp.NewOrganizationController(orgQueryHandler, orgCommandHandler)
	memberCommandHandler := orgHandlers.NewMemberCommandHandler(membershipRepo, orgRepo)
	memberController := orgHttp.NewMemberController(orgQueryHandler, orgCommandHandler, memberCommandHandler)

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
		DB:               db,
		Redis:            redisClient,
		Log:              logger,
		Config:           cfg,
		JWTService:       jwtService,
		PasswordService:  passwordService,
		AuthService:      authSvc,
		AuthController:   authController,
		AuthGrpcServer:   authGrpcServer,
		UserController:   userController,
		OrgController:    orgController,
		MemberController: memberController,
	}, nil
}
