package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	bloom "github.com/bits-and-blooms/bloom/v3"
)

func main() {
	var i = 1
	filter := bloom.NewWithEstimates(850000000, 0.00001)
	// 850 million entries, 0.00001% false positive rate
	file, err := os.Open("goodhash.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var myString = scanner.Text()
		myString = strings.ToLower(myString)

		filter.AddString(myString)
		if i%1000000 == 0 {
			fmt.Println(i/1000000, " million entries added")
		}
		i++
		if myString == "5b4c6335673a75f13ed948e848f00840" {
			fmt.Println("Found it!")
		}

	}
	file.Close()

	if filter.TestString("5b4c6335673a75f13ed948e848f00840") {
		fmt.Println("Love is in the filter")
	} else {
		fmt.Println("Love is not in the filter")
	}

	file, err = os.Create("bloom.bin")
	if err != nil {
		fmt.Println(err)
	}
	filter.WriteTo(file)
	file.Close()
}
