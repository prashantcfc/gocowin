package log

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	stdlog "log"
	"os"
)

func GetLoggerWithName(name string) logr.Logger {
	stdr.SetVerbosity(1)
	log := stdr.NewWithOptions(stdlog.New(os.Stderr, "", stdlog.LstdFlags), stdr.Options{LogCaller: stdr.All})
	log = log.WithName(name)

	return log
}