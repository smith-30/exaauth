package logger

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/emonuh/zap-error-dispatcher/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(c Config) (*Logger, error) {
	loggerConf, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	var zapConf config.ErrorDispatcherConfig
	if err := json.Unmarshal(loggerConf, &zapConf); err != nil {
		return nil, err
	}

	zapConf.EncoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.UTC().Format(timeFormat))
	}
	zapConf.EncoderConfig.EncodeDuration = func(d time.Duration, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(fmt.Sprintf("%.3fms", float64(d)/float64(time.Millisecond)))
	}

	l, err := zapConf.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{l}, nil
}

func (a *Logger) Panic(err interface{}) {
	a.Error("", zap.Any("panic_by", err), zap.String("stack", string(debug.Stack())))
}
