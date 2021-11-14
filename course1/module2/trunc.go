package main

import "fmt"

func main() {
	var userInput float32
	fmt.Printf("Enter a floating point number:")
	fmt.Scan(&userInput)
	fmt.Printf("The truncated version of the number you entered is %d", int(userInput))
}
