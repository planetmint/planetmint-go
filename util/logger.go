package util

import sdk "github.com/cosmos/cosmos-sdk/types"

type AppLogger struct {
}

var globalApplicationLoggerTag string

func init() {
	// Initialize the package-level variable
	globalApplicationLoggerTag = "[app]"
}

func GetAppLogger() *AppLogger {
	return &AppLogger{}
}

func (logger *AppLogger) Info(ctx sdk.Context, msg string, keyvals ...interface{}) {
	ctx.Logger().Info(globalApplicationLoggerTag+msg, keyvals)

}

func (logger *AppLogger) Debug(ctx sdk.Context, msg string, keyvals ...interface{}) {
	ctx.Logger().Debug(globalApplicationLoggerTag+msg, keyvals)

}

func (logger *AppLogger) Error(ctx sdk.Context, msg string, keyvals ...interface{}) {
	ctx.Logger().Error(globalApplicationLoggerTag+msg, keyvals)
}
