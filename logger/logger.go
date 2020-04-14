package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

const (
	//only using in developping, and please change to false before releasing binary executives
	isDEBUG = true
)

type thelogger struct {
	logger *log.Logger
}

//GLogger -- global logger.
var GLogger = &thelogger{logger: log.New(os.Stdout, "CLUSTERSHELL-", log.LstdFlags)}


//SetPrefix --
func SetPrefix(prefix string) {
	GLogger.logger.SetPrefix(prefix)
}

//Info --
func Info(params ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	GLogger.logger.Print(fmt.Sprintf("[Info] in %s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line), params)
	//GLogger.logger.Print("Info -- ", params)
}

//Warning --
func Warning(params ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	GLogger.logger.Print(fmt.Sprintf("[Warning] in %s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line), params)
	//GLogger.logger.Print("Warning -- ", params)
}

//Error --
func Error(params ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	GLogger.logger.Print(fmt.Sprintf("[Error] in %s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line), params)
	//GLogger.logger.Print("Error -- ", params)
}

//Debug --
func Debug(params ...interface{}) {
	if isDEBUG {
		pc, fn, line, _ := runtime.Caller(1)
		GLogger.logger.Print(fmt.Sprintf("[Debug] in %s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line), params)
		//	GLogger.logger.Print("Debug -- ", params)
	}
}
