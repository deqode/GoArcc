package logger

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	// Log is global logger
	Log *zap.Logger

	// customTimeFormat is custom Time format
	customTimeFormat string
)

type Config struct {
	LogLevel    zapcore.Level
	Development bool
}

// customTimeEncoder encode Time to our custom format
// This example how we can customize zap default functionality
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(customTimeFormat))
}

// Init initializes log by input parameters
// lvl - global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
// timeFormat - custom time format for logger of empty string to use default
func Init(config Config) (*zap.Logger, error) {

	if Log != nil {
		Log.Fatal("Logger already initialized once, No need to do it multiple times")
		return nil, errors.New("logger already initialized once")
	}

	// First, define our level-handling logic.
	globalLevel := config.LogLevel
	logTimeFormat := ""

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	// It is useful for Kubernetes deployment.
	// Kubernetes interprets os.Stdout log items as INFO and os.Stderr log items
	// as ERROR by default.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= globalLevel && lvl < zapcore.ErrorLevel
	})
	consoleInfos := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Configure console output.
	var useCustomTimeFormat bool

	var ecfg zapcore.EncoderConfig
	if config.Development {
		ecfg = zap.NewDevelopmentEncoderConfig()
	} else {
		ecfg = zap.NewProductionEncoderConfig()
	}

	if len(logTimeFormat) > 0 {
		customTimeFormat = logTimeFormat
		ecfg.EncodeTime = customTimeEncoder
		useCustomTimeFormat = true
	}
	consoleEncoder := zapcore.NewJSONEncoder(ecfg)

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority),
	)
	// From a zapcore.Core, it's easy to construct a Logger.
	Log = zap.New(core)
	zap.ReplaceGlobals(Log)
	zap.RedirectStdLog(Log)

	if !useCustomTimeFormat {
		Log.Warn("time format for logger is not provided - use zap default")
	}

	return Log, nil
}
