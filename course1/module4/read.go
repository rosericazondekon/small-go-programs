package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Variable definition
	var fileName string
	type Person struct{ firstName, lastName string }
	names := make([]Person, 0)

	// Prompt user to enter file name
	fmt.Printf("Enter your file name: ")
	nameScanner := bufio.NewScanner(os.Stdin)
	if nameScanner.Scan() {
		fileName = strings.TrimSpace(nameScanner.Text())
	}

	// Processing and Output
	dat, err := ioutil.ReadFile(fileName)

	if err == nil {
		data := strings.Split(string(dat), "\n")
		for i := 0; i < len(data); i++ {
			s := strings.Fields(data[i])
			p := Person{firstName: s[0], lastName: s[1]}
			names = append(names, p)
		}

		for i := 0; i < len(names); i++ {
			fmt.Printf("First Name: %s, Last Name: %s\n", names[i].firstName, names[i].lastName)
		}

	} else {
		fmt.Println("Please, enter a valid file name!")
	}
}
