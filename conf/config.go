package conf

var Config = &AppConfig{}

type AppConfig struct {
	AppName     string
	AppPath     string
	AppPort     int
	AppVersion  string
	Timezone    string
	RunMode     string
	HttpLimiter float64 // 每秒最大访问量

	MySQL map[string]MySQLConfig
	Redis map[string]RedisConfig

	Grpc    grpcConfig
	Log     LogConfig
	Warnbot WarnbotConfig
}

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
	TablePre string
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
	Dir        string
	Level      string
	ToFile     bool
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}
type WarnbotConfig struct {
	Wx string
}
