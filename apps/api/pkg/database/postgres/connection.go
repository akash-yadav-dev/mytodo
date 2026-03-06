package postgres

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
