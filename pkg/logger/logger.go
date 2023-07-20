package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type (
	Config struct {
		Logger
	}

	Logger struct {
		Level int8 `koanf:"level" validate:"required"`
	}
)

func NewLogger(cfg Config) *zap.SugaredLogger {
	atom := zap.NewAtomicLevel()
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		atom,
	)
	logger := zap.New(zapCore)

	l := logger.Sugar()
	atom.SetLevel(zapcore.Level(cfg.Logger.Level))

	return l
}

func DeferLogger(l *zap.SugaredLogger) {
	_ = l.Sync()
	return
}
