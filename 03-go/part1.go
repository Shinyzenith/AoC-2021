package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	positiveBits := make([]int, 12)
	negativeBits := make([]int, 12)

	for _, bin := range GetFileData() {
		for i, char := range bin {
			if strings.EqualFold(string(char), "1") {
				positiveBits[i]++
			} else {
				negativeBits[i]++
			}
		}
	}

	most := ""
	least := ""
	for i := 0; i < len(positiveBits); i++ {
		if positiveBits[i] > negativeBits[i] {
			most += "1"
			least += "0"
		} else {
			most += "0"
			least += "1"
		}
	}

	first, err := strconv.ParseInt(most, 2, 64)
	Must(err)
	second, err := strconv.ParseInt(least, 2, 64)
	Must(err)

	fmt.Println(first * second)
}

func GetFileData() []string { // Read the file into an array of ints
	content, err := os.ReadFile("./data")
	Must(err)
	return strings.Split(string(content), "\n")
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
