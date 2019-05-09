package logger

import (
	"go.uber.org/zap"
)

// NewTmpLogger create SugaredLogger.
// It is used to keep a log at startup. For example before loading config file.
func NewTmpLogger() (*zap.SugaredLogger, error) {
	c := newDefaultConfig()
	l, err := NewLogger(c)
	if err != nil {
		return nil, err
	}

	return l.Sugar(), nil
}
