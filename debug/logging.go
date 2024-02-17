package debug

import (
	"os"

	"github.com/rs/zerolog"
)

func SetupLogger() *zerolog.Logger {
	// set up logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if os.Getenv("DOPEWARS_DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log := zerolog.New(os.Stdout).With().
		Timestamp().Caller().Logger().Output(
		// nolint:exhaustruct // not needed here
		zerolog.ConsoleWriter{Out: os.Stderr})
	return &log
}
