package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var userInput string
	fmt.Printf("Enter a sequence of integers separated by a space: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = strings.TrimSpace(scanner.Text())
	}
	data := strings.Split(userInput, " ")
	var sli []int64
	for i, _ := range data {
		numb, e := strconv.ParseInt(data[i], 0, 64)
		if e == nil {
			sli = append(sli, numb)
		}
	}
	fmt.Printf("UnSorted sequence: %d\n\n", sli)
	BubbleSort(sli)
	fmt.Printf("Sorted sequence (from least to greatest): %d\n", sli)
}

func Swap(sli []int64, i int) {
	prev := sli[i-1]
	sli[i-1] = sli[i]
	sli[i] = prev
}

func BubbleSort(sli []int64) {
	n := len(sli)
	for n > 1 {
		newn := 0
		for i := 1; i < n; i++ {
			if sli[i-1] > sli[i] {
				Swap(sli, i)
				newn = i
			}
		}
		n = newn
	}
}