package main

import (
	"backup-automation/cli"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {

	logsPath := "../Logs/logs.log"

	file, err := os.OpenFile(logsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)

		logrus.AddHook(lfshook.NewHook(lfshook.PathMap{
			logrus.InfoLevel:  "../Logs/info.log",
			logrus.ErrorLevel: "../Logs/error.log",
			logrus.WarnLevel:  "../Logs/warn.log",
			logrus.FatalLevel: "../Logs/fatal.log",
			logrus.PanicLevel: "../Logs/panic.log",
			logrus.DebugLevel: "../Logs/debug.log",
			logrus.TraceLevel: "../Logs/trace.log",
		}, &logrus.JSONFormatter{}))
	}
}

func main() {
	cli.CLI()
}
