package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string

	fmt.Printf("Enter a name: ")
	nameScanner := bufio.NewScanner(os.Stdin)
	if nameScanner.Scan() {
		name = nameScanner.Text()
	}

	fmt.Printf("Enter an address: ")
	adScanner := bufio.NewScanner(os.Stdin)
	if adScanner.Scan() {
		address = adScanner.Text()
	}

	m := map[string]string{"name": name, "address": address}
	infoJSON, err := json.Marshal(m)

	if err == nil {
		fmt.Printf("%s\n", infoJSON)
	}
}
