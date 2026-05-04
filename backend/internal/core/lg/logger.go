package lg

import (
	"fmt"
	"time"
)

type Color string

const (
	Reset   Color = "\033[0m"
	End     Color = "\033[0m\n"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Purple  Color = "\033[35m"
	Cyan    Color = "\033[36m"
	Gray    Color = "\033[90m"
	UBlack  Color = "\033[4;30m"
	URed    Color = "\033[4;31m"
	UGreen  Color = "\033[4;32m"
	UYellow Color = "\033[4;33m"
	UBlue   Color = "\033[4;34m"
	UPurple Color = "\033[4;35m"
	UCyan   Color = "\033[4;36m"
	UWhite  Color = "\033[4;37m"
)

type LoggerFunc func(msg string)

type Logger struct {
	Info  LoggerFunc
	Warn  LoggerFunc
	Error LoggerFunc
}

func logMessage(levelColor, msgColor Color, level, underLineColor Color, msg string) {
	// Combina nivel y mensaje en un único formato
	date := time.Now().Format("2006-01-02 15:04:05 (utc)")
	format := fmt.Sprintf("%s%s - %s:%s %s%s%s%s",
		levelColor,
		date,
		level,
		Reset,
		underLineColor,
		msgColor,
		msg,
		End,
	)
	fmt.Println(format)
}

// recibe multiple args for logging
func Info(msg string)  { logMessage(Blue, Green, "[Info]", UGreen, msg) }
func Warn(msg string)  { logMessage(Yellow, Yellow, "[Warn]", UYellow, msg) }
func Error(msg string) { logMessage(Red, Blue, "[Error]", URed, msg) }
