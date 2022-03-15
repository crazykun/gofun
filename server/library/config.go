package library

type Config struct {
	AppName  string
	LogLevel string

	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}
