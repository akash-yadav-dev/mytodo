package bootstrap

type Config struct {
	// Server

	ServerPort string

	//Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	//Redis
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	//jwt
	JWTSecret string
	JWTExpiry int
}
