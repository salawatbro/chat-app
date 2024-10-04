package utils

import (
	"github.com/salawatbro/chat-app/pkg/constants"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger app logger by uber zap
var Logger *zap.Logger

func ZapLogger(env string) {
	// encoder config per env
	var logLevel zapcore.LevelEnabler
	var encoderConfig zapcore.EncoderConfig
	if strings.ToLower(env) == "production" {
		encoderConfig = zap.NewProductionEncoderConfig()
		logLevel = zapcore.ErrorLevel
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		logLevel = zapcore.DebugLevel
	}

	encoderConfig.EncodeCaller = nil // hide caller (filename)
	encoderConfig.EncodeLevel = nil  // hide log level
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(constants.TimestampFormat)
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), logLevel)
	logger := zap.New(core, zap.AddCaller())
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
	Logger = logger
}
