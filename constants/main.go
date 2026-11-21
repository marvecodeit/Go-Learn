package main

import "fmt"

const (
	HOST = "127.0.0.0.1"
	PORT = ":8080"
	USER = "root"
)

var (
	isRunning bool
)

func main() {
	AppName := "Go"
	fmt.Println(AppName)

	const pi float64 = 3.1415926
	fmt.Println(pi)

	const rate float32 = 5.7
	fmt.Println(rate)
}
