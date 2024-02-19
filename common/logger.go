package common

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func GetLogger() *log.Logger {

	var Log *log.Logger
	Log = log.New()
	Log.SetFormatter(&log.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	return Log
}
