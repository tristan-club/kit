package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tristan-club/kit/config"
	"os"
	"runtime"
	"strconv"
	"time"
)

var logger *zerolog.ConsoleWriter

const (
	traceId = "traceid"
)

func init() {

	if config.EnvIsDev() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.TimeFieldFormat = time.RFC3339Nano
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if config.UseConsoleWrite() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}
}

func Error() *zerolog.Event {

	log.With().Caller()

	_, file, line, ok := runtime.Caller(1)
	e := log.Error()
	if ok {
		e = e.Str("line", file+":"+strconv.Itoa(line))
	}
	return e
}

func Debug() *zerolog.Event {
	_, file, line, ok := runtime.Caller(1)
	e := log.Debug()
	if ok {
		e = e.Str("line", file+":"+strconv.Itoa(line))
	}
	return e
}

func Warn() *zerolog.Event {
	_, file, line, ok := runtime.Caller(1)
	e := log.Warn()
	if ok {
		e = e.Str("line", file+":"+strconv.Itoa(line))
	}
	return e
}

func Info() *zerolog.Event {
	_, file, line, ok := runtime.Caller(1)
	e := log.Info()
	if ok {
		e = e.Str("line", file+":"+strconv.Itoa(line))
	}
	return e
}
