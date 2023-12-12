package util

import (
	"sync"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AppLogger struct {
	testingLogger *testing.T
}

var (
	globalApplicationLoggerTag string
	appLogger                  *AppLogger
	initAppLogger              sync.Once
)

func init() {
	// Initialize the package-level variable
	globalApplicationLoggerTag = "[app] "
}

func GetAppLogger() *AppLogger {
	initAppLogger.Do(func() {
		appLogger = &AppLogger{
			testingLogger: nil,
		}
	})
	return appLogger
}

func (logger *AppLogger) SetTestingLogger(testingLogger *testing.T) *AppLogger {
	logger.testingLogger = testingLogger
	return logger
}

func (logger *AppLogger) testingLog(msg string, keyvals ...interface{}) {
	if logger.testingLogger == nil {
		return
	}
	logger.testingLogger.Log(msg, keyvals)
}

func (logger *AppLogger) Info(ctx sdk.Context, msg string, keyvals ...interface{}) {
	logger.testingLog(globalApplicationLoggerTag+msg, keyvals...)
	ctx.Logger().Info(globalApplicationLoggerTag+msg, keyvals...)
}

func (logger *AppLogger) Debug(ctx sdk.Context, msg string, keyvals ...interface{}) {
	logger.testingLog(globalApplicationLoggerTag+msg, keyvals...)
	ctx.Logger().Debug(globalApplicationLoggerTag+msg, keyvals...)
}

func (logger *AppLogger) Error(ctx sdk.Context, msg string, keyvals ...interface{}) {
	logger.testingLog(globalApplicationLoggerTag+msg, keyvals...)
	ctx.Logger().Error(globalApplicationLoggerTag+msg, keyvals...)
}
