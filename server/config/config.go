package config

type Config struct {
	RDB    RDB
	Server Server
	Logger Logger
}

type RDB struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

type Server struct {
	Host      string
	Port      string
	JWTSecret string
}

type Logger struct {
	Development bool
	Level       string

	// "console" or "json"
	Encoding            string
	OutputPaths         []string
	AppErrorOutputPaths []string
	ErrorOutputPaths    []string
	EncoderConfig
}

type EncoderConfig struct {
	MessageKey    string
	LevelKey      string
	TimeKey       string
	NameKey       string
	CallerKey     string
	StacktraceKey string

	// "capital", "capitalColor" or "color". default is lowercase
	LevelEncoder  string
	CallerEncoder string
}

func NewDefaultConfig() *Config {
	return &Config{}
}
