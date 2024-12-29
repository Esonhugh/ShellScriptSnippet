package log

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var GlobalLogLevel = LogLevelInfo

const (
	LogLevelTrace = 2
	LogLevelDebug = 1
	LogLevelInfo  = 0
)

// Init func is a function to init logrus with specific log level
func Init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(logFormat())
	log.SetLevel(logLevel(GlobalLogLevel))
}

// logLevel search level strings return correct Level
func logLevel(level int) log.Level {
	switch level {
	case LogLevelTrace:
		return log.TraceLevel
	case LogLevelDebug:
		return log.DebugLevel
	case LogLevelInfo:
		return log.InfoLevel
	default:
		return log.InfoLevel
	}
}

// logFormat sets log format by using prefixed "x-cray/logrus-prefixed-formatter"
func logFormat() log.Formatter {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.SetColorScheme(&prefixed.ColorScheme{
		PrefixStyle:    "blue+b",
		TimestampStyle: "white+h",
	})
	return formatter
}
