package main

import (
	"fmt"
	"strings"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
}

func (clog ConsoleLogger) Log(message string) {
	fmt.Println("[CONSOLE]", message)
}

type FileLogger struct {
}

func (flog FileLogger) Log(message string) {
	fmt.Println("[FILE]", message)
}

func ProcessData(logger Logger, data string) {
	result := strings.ToUpper(data)
	logger.Log(result)
}

func main() {
	clog := ConsoleLogger{}
	flog := FileLogger{}
	loggers := []Logger{clog, flog}
	for _, l := range loggers {
		ProcessData(l, "Someting")
	}
}
