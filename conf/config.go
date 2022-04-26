package conf

var Config = &AppConfig{}

type AppConfig struct {
	AppName    string
	AppPath    string
	AppPort    int
	AppVersion string
	Timezone   string
	RunMode    string

	MySQL map[string]MySQLConfig
	Redis map[string]RedisConfig

	Grpc grpcConfig
	Log  LogConfig
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
	MaxNum int
	MinNum int
}

type grpcConfig struct {
	Host string
	Port int
}

type LogConfig struct {
	Dir   string
	Level string
}
