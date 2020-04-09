package logger

import (
	"log"
	"os"
)

const (
	//only using in developping, and please change to false before releasing binary executives
	isDEBUG = true
)

type thelogger struct {
	logger *log.Logger
}

//GLogger -- global logger.
var GLogger = &thelogger{logger: log.New(os.Stdout, "CLUSTERSHELL-", log.Lshortfile)}

//Info --
func Info(params ...interface{}) {
	GLogger.logger.Print("Info -- ", params)
}

//Warning --
func Warning(params ...interface{}) {
	GLogger.logger.Print("Warning -- ", params)
}

//Error --
func Error(params ...interface{}) {
	GLogger.logger.Print("Error -- ", params)
}

//Debug --
func Debug(params ...interface{}) {
	if isDEBUG {
		GLogger.logger.Print("Debug -- ", params)
	}
}
