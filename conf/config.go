package conf

var Config = &AppConfig{}
var Db = &MySQLConfig{}
var Redis = &RedisConfig{}

type AppConfig struct {
	AppName    string
	AppPath    string
	AppPort    int
	AppVersion string
	Timezone   string
	LogLevel   string
	RunMode    string

	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Charset  string
	MaxMum   int
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
