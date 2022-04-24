package conf

var Config = &AppConfig{}

type AppConfig struct {
	AppName    string
	AppPath    string
	AppPort    int
	AppVersion string
	Timezone   string
	LogLevel   string
	RunMode    string

	MySQL map[string]MySQLConfig
	Redis map[string]RedisConfig

	Grpc grpcConfig
}

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Charset  string
	MaxNum   int
	MinNum   int
}

type RedisConfig struct {
	Host   string
	Port   int
	Auth   string
	Db     int
	MaxMum int
	MinNum int
}

type grpcConfig struct {
	Host string
	Port int
}
