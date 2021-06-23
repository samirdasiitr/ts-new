package utils

import (
	"log"
	"os"
)

// Logger struct.
type logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
	fatalLogger *log.Logger
	debugLogger *log.Logger
}

// Info print.
func (_lgr *logger) INFO(format string, a ...interface{}) {
	_lgr.infoLogger.Printf(format, a...)
}

// Error print.
func (_lgr *logger) ERROR(format string, a ...interface{}) {
	_lgr.errorLogger.Printf(format, a...)
}

// Warn print.
func (_lgr *logger) WARN(format string, a ...interface{}) {
	_lgr.warnLogger.Printf(format, a...)
}

// Fatal print.
func (_lgr *logger) FATAL(format string, a ...interface{}) {
	_lgr.fatalLogger.Fatalf(format, a...)
}

// Debug print.
func (_lgr *logger) DEBUG(format string, a ...interface{}) {
	_lgr.fatalLogger.Fatalf(format, a...)
}

func (_logger *logger) Init() {
	file, err := os.OpenFile("logs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	_logger.infoLogger = log.New(file, "INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	_logger.errorLogger = log.New(file, "ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	_logger.warnLogger = log.New(file, "WARN: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	_logger.fatalLogger = log.New(file, "FATAL: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	_logger.fatalLogger = log.New(file, "DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	_logger.infoLogger.Println("New Logging Session Started")
}

var Log = &logger{}
