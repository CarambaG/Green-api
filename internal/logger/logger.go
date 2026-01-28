package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLog  = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	errorLog = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	warnLog  = log.New(os.Stdout, "[WARN] ", log.LstdFlags)
)

func Info(format string, args ...interface{}) {
	infoLog.Println(fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	errorLog.Println(fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	warnLog.Println(fmt.Sprintf(format, args...))
}
