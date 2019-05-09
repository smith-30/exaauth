package logger

const (
	timeFormat = "2006-01-02T15:04:05.000000Z07:00"
)

type (
	Config struct {
		Development         bool     `json:"development"`
		Level               string   `json:"level"`
		Encoding            string   `json:"encoding"`
		OutputPaths         []string `json:"outputPaths"`
		AppErrorOutputPaths []string `json:"errorDispatcherPaths"` // Adjust to library
		ErrorOutputPaths    []string `json:"errorOutputPaths"`
		EncoderConfig       EncoderConfig
	}

	EncoderConfig struct {
		MessageKey    string `json:"messageKey"`
		LevelKey      string `json:"levelKey"`
		TimeKey       string `json:"timeKey"`
		NameKey       string `json:"nameKey"`
		CallerKey     string `json:"callerKey"`
		StacktraceKey string `json:"stacktraceKey"`
		LevelEncoder  string `json:"levelEncoder"`
		CallerEncoder string `json:"callerEncoder"`
	}
)

func newDefaultConfig() Config {
	return Config{
		Level:               "info",
		Encoding:            "json",
		OutputPaths:         []string{"stdout"},
		AppErrorOutputPaths: []string{"stderr"},
		ErrorOutputPaths:    []string{"stderr"},
		EncoderConfig: EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "ts",
			NameKey:       "name",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			LevelEncoder:  "",
			CallerEncoder: "short",
		},
	}
}
