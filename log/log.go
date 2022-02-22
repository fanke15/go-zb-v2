package log

import (
	"github.com/dwdcth/consoleEx"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	out := consoleEx.ConsoleWriterEx{Out: colorable.NewColorableStdout()}
	Log = zerolog.New(out).With().Caller().Timestamp().Logger()

	//Log = log.With().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "01-02 15:04:05", NoColor: false})
	//Log = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "01-02 15:04:05", NoColor: false})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
