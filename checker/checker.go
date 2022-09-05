package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	bloom "github.com/bits-and-blooms/bloom/v3"
)

func main() {

	newfilter := bloom.NewWithEstimates(850000000, 0.00001)

	file, err := os.Open("./bloom.bin")

	if err != nil {
		fmt.Println(err)
	}

	newfilter.ReadFrom(file)

	file.Close()

	dump, err := os.Open("../output.ntds")
	defer dump.Close()
	if err != nil {
		fmt.Println(err)
	}
	var badOnes []string
	scanner := bufio.NewScanner(dump)
	for scanner.Scan() {
		var myString = scanner.Text()
		var hash = strings.Split(myString, ":")[3]
		var user = strings.Split(myString, ":")[0]

		if newfilter.TestString(hash) {
			badOnes = append(badOnes, user)
		}

	}
	// print bad ones to file
	f, err := os.Create("badOnes.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for _, item := range badOnes {
		_, err := f.WriteString(item + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}

}
