package main

import (
	"io"
	"os"
	"fmt"
	logger "github.com/sirupsen/logrus"
)

var mainLogger *logger.Entry
var webLogger *logger.Entry

func initLogger() {
	logFile, err := os.OpenFile("latest.log", os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("ERROR: Failed to create log file. Application quitting now")
		os.Exit(2)
	}
	logger.SetLevel(logger.InfoLevel)
	mw := io.MultiWriter(os.Stdout, logFile)
	logger.SetOutput(mw)
	logger.SetFormatter(&logger.JSONFormatter{})
	// Set loggers
	mainLogger = logger.WithFields(logger.Fields{
		"application": "server-ctl.main",
	})
	webLogger = logger.WithFields(logger.Fields{
		"application": "server-ctl.web",
	})
}
