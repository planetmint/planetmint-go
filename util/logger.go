package util

import sdk "github.com/cosmos/cosmos-sdk/types"

type AppLogger struct {
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
		}
	})
	return appLogger
}

func (logger *AppLogger) Info(ctx sdk.Context, msg string, keyvals ...interface{}) {
	ctx.Logger().Info(globalApplicationLoggerTag+msg, keyvals...)
}

func (logger *AppLogger) Debug(ctx sdk.Context, msg string, keyvals ...interface{}) {
	ctx.Logger().Debug(globalApplicationLoggerTag+msg, keyvals...)
}

func (logger *AppLogger) Error(ctx sdk.Context, msg string, keyvals ...interface{}) {
	ctx.Logger().Error(globalApplicationLoggerTag+msg, keyvals...)
}
