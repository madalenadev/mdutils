package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// These consts are refering levels/severity from GCP
const (
	DEFAULT   = "Default"
	DEBUG     = "Debug"
	INFO      = "Info"
	NOTICE    = "Notice"
	WARNING   = "Warning"
	ERROR     = "Error"
	ALERT     = "Alert"
	EMERGENCY = "Emergency"
)

var (
	log              *logrus.Logger
	useSeverityField bool
	fields           *logrus.Fields
)

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
	useSeverityField = true
}

// SetLevel altera o level do logger
func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	log.SetLevel(lvl)
}

// GetLevel recupera o level do logger
func GetLevel() logrus.Level {
	return log.GetLevel()
}

// DisableSeverityField it will do disable severity fields on log
func DisableSeverityField() {
	useSeverityField = false
}

// setFieldSeverity will set severity field when it was availble
func setFieldSeverity(lvl string) {
	if useSeverityField {
		fields = &logrus.Fields{"severity": lvl}
	} else {
		fields = &logrus.Fields{}
	}
}

// Warning show detail of log
func Warning(args ...interface{}) {
	setFieldSeverity(WARNING)
	log.WithFields(*fields).Warning(args...)
}

func Error(args ...interface{}) {
	setFieldSeverity(ERROR)
	log.WithFields(*fields).Error(args...)
}

func Info(args ...interface{}) {
	setFieldSeverity(INFO)
	log.WithFields(*fields).Info(args...)
}

func Debug(args ...interface{}) {
	setFieldSeverity(DEBUG)
	log.WithFields(*fields).Debug(args...)
}

func Trace(args ...interface{}) {
	setFieldSeverity(DEBUG)
	log.WithFields(*fields).Trace(args...)
}

func Fatal(args ...interface{}) {
	setFieldSeverity(EMERGENCY)
	log.WithFields(*fields).Panic(args...)
}
