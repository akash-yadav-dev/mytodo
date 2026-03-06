package bootstrap

type App struct {
	container *Container
	logger    Logger
}

func NewApp(container *Container, logger Logger) *App {
	return &App{
		container: container,
		logger:    logger,
	}
}

func (a *App) GetAuthHandler() *auth.AuthHandler {
	return a.container.AuthHandler
}

func (a *App) GetUserHandler() *users.UserHandler {
	return a.container.UserHandler
}

func (a *App) GetProjectHandler() *projects.ProjectHandler {
	return a.container.ProjectHandler
}

func (a *App) GetIssueHandler() *issues.IssueHandler {
	return a.container.IssueHandler
}

func (a *App) GetLogger() Logger {
	return a.logger
}
