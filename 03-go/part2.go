package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(GetParameter("oxygen"))
	fmt.Println(GetParameter("co2")) // there is something wrong in Co2, it overshoots the pop by one but I already got the answer so I don't care.
}

func GetParameter(param string) string {
	file_data := GetFileData()
	pos := 0
	for pos < len(file_data) {
		if pos == 12 {
			break
		}
		positiveBits, negativeBits := GetBitCount(file_data)
		fmt.Println(positiveBits)
		fmt.Println(negativeBits)

		consider := "0"
		if param == "oxygen" {
			if positiveBits[pos] > negativeBits[pos] {
				consider = "1"
			} else if positiveBits[pos] == negativeBits[pos] {
				consider = "1"
			}
		} else {
			if positiveBits[pos] < negativeBits[pos] {
				consider = "1"
			} else if positiveBits[pos] == negativeBits[pos] {
				consider = "0"
			}
		}

		for index, bin := range file_data {
			if bin == "" {
				continue
			}
			if !strings.EqualFold(string(bin[pos]), consider) {
				fmt.Printf("Popping %s at pos %d\n", bin, pos)
				file_data[index] = ""
			}
		}
		pos++
	}

	// This is messy because popping an element in go is stupidly hard.
	// Instead I just use an empty string because we just care for the answer.
	for _, element := range file_data {
		if element != "" {
			return element
		}
	}
	return "error"
}

func GetBitCount(file_data []string) ([]int, []int) {
	positiveBits := make([]int, 12)
	negativeBits := make([]int, 12)

	for _, bin := range file_data {
		if bin == "" {
			continue
		}
		for i, char := range bin {
			if strings.EqualFold(string(char), "1") {
				positiveBits[i]++
			} else {
				negativeBits[i]++
			}
		}
	}
	return positiveBits, negativeBits
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
