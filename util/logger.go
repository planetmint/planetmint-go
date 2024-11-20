package util

import (
	"fmt"
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
	syncTestingLog             sync.RWMutex
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
	syncTestingLog.Lock()
	logger.testingLogger = testingLogger
	syncTestingLog.Unlock()
	return logger
}

func format(msg string, keyvals ...interface{}) string {
	if len(keyvals) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, keyvals...)
}

func (logger *AppLogger) testingLog(msg string, keyvals ...interface{}) {
	syncTestingLog.RLock()
	defer syncTestingLog.RUnlock()
	if logger.testingLogger == nil {
		return
	}
	msg = format(msg, keyvals...)
	logger.testingLogger.Log(msg)
}

func (logger *AppLogger) Info(ctx sdk.Context, msg string, keyvals ...interface{}) {
	msg = format(msg, keyvals...)
	logger.testingLog(globalApplicationLoggerTag + msg)
	ctx.Logger().Info(globalApplicationLoggerTag + msg)
}

func (logger *AppLogger) Debug(ctx sdk.Context, msg string, keyvals ...interface{}) {
	msg = format(msg, keyvals...)
	logger.testingLog(globalApplicationLoggerTag + msg)
	ctx.Logger().Debug(globalApplicationLoggerTag + msg)
}

func (logger *AppLogger) Error(ctx sdk.Context, err error, msg string, keyvals ...interface{}) {
	msg = format(msg, keyvals...)
	logger.testingLog(globalApplicationLoggerTag + msg + ": " + err.Error())
	ctx.Logger().Error(globalApplicationLoggerTag + msg + ": " + err.Error())
}
