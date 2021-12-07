package log

import (
	"github.com/sirupsen/logrus"
	"os"
)
var Log *logrus.Logger
func InitLogrus()  {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)                     //default
	//log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout
	Log = log
	log.Debug("init Log-rus Success")

}