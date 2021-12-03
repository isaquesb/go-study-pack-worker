package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"worker/core/port"
)

type zeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger() port.Logger {
	zerolog.CallerFieldName = "trace"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("ENVIRONMENT") == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return &zeroLogger{log.With().
		Str("app", os.Getenv("APP")).
		Str("version", os.Getenv("VERSION")).
		Logger(),
	}
}

func (l zeroLogger) Log(Message string, Context ...string) {
	l.logger.Trace().Strs("context", Context).Msg(Message)
}

func (l zeroLogger) Debug(Message string, Context ...string) {
	l.logger.Debug().Strs("context", Context).Msg(Message)
}

func (l zeroLogger) Notice(Message string, Context ...string) {
	l.logger.Info().Strs("context", Context).Msg(Message)
}

func (l zeroLogger) Info(Message string, Context ...string) {
	l.logger.Info().Strs("context", Context).Msg(Message)
}

func (l zeroLogger) Warning(Message string, Context ...string) {
	l.logger.Warn().Strs("context", Context).Msg(Message)
}

func (l zeroLogger) Error(Error error, Context ...string) {
	l.logger.Error().Strs("context", Context).Caller().Msg(Error.Error())
}

func (l zeroLogger) Critical(Error error, Context ...string) {
	l.logger.Fatal().Strs("context", Context).Caller().Msg(Error.Error())
}

func (l zeroLogger) Alert(Error error, Context ...string) {
	l.logger.Panic().Strs("context", Context).Caller().Msg(Error.Error())
}

func (l zeroLogger) Emergency(Error error, Context ...string) {
	l.logger.Panic().Strs("context", Context).Caller().Msg(Error.Error())
}
