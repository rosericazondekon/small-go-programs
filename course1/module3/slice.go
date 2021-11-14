package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	x := make([]int, 0, 3)
	for i := 0; ; i++ {
		var input string
		fmt.Printf("Enter an integer (Enter 'X' to Quit): ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = strings.TrimSpace(scanner.Text())
		}
		if input == "X" {
			fmt.Println(x)
			break
		}
		val, err := strconv.Atoi(input)
		if err == nil {
			x = append(x, val)
			sort.Ints(x)
			fmt.Println(x)
		} else {
			fmt.Println(x)
		}
	}
}
