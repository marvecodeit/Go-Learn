package main

import "fmt"

type LoggerLevel int

const (
	LevelTrace LoggerLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelName = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LoggerLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unkonwn"
	}

	return levelName[l]
}

func printLogLevel(level LoggerLevel) {
	fmt.Printf("Log Level: %d %s\n", level, level.String())
}
func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(10)
}
