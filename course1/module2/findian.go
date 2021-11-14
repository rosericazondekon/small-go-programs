package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput string
	fmt.Printf("Enter a string: ")
	// fmt.Scanln(&userInput)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}

	userInput = strings.ToLower(userInput)
	bool_status := strings.HasPrefix(userInput, "i") && strings.HasSuffix(userInput, "n") && strings.Contains(userInput, "a")
	if bool_status == true {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
