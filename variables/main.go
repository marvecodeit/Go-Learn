package main

import "fmt"

func main() {
	var greeting string
	greeting = "Hello World"

	fmt.Println(greeting)

	var count int
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = false || true
	fmt.Println(isRunning)

	firstname, lastname, score := "Marvellous", "Onyema", 8
	fmt.Println(firstname, lastname, score)

	var surname, othername string
	surname = "Marvellous"
	othername = "Ugochukwu"
	fmt.Println(surname, othername)

	email, age := "marvecodeit.dev", 8
	fmt.Println(email, age)

	var year = 2026
	fmt.Println(year)
}
