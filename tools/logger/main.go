package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	successColor = 32
	warningColor = 33
	errorColor   = 31
	infoColor    = 36
	debugColor   = 35
)

func log(color int, messages []interface{}) {

	datetime := time.Now().Format("2006-01-02 15:04:05")

	pc, _, _, _ := runtime.Caller(2)
	function := runtime.FuncForPC(pc).Name()

	fmt.Printf("\033[%dm●\033[0m | %v | %v | ", color, datetime, function)

	for _, message := range messages {
		fmt.Printf("%v ", message)
	}

	fmt.Println()

}

func Success(messages ...interface{}) {
	log(successColor, messages)
}

func Warning(messages ...interface{}) {
	log(warningColor, messages)
}

func Error(messages ...interface{}) {
	log(errorColor, messages)
	os.Exit(1)
}

func Info(messages ...interface{}) {
	log(infoColor, messages)
}

func Debug(messages ...interface{}) {
	log(debugColor, messages)
}
