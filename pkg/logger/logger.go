package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var log *zerolog.Logger = nil

func Init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()

	log = &logger
}

func Infof(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatal().Msgf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debug().Msgf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}
