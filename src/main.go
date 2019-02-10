package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Check if file path is given
	if len(os.Args) != 2 {
		fmt.Println("You need to give the program a file.")
		fmt.Println("Usage:", os.Args[0], "[FILE_PATH]")
		os.Exit(1)
	}

	// Read the provided file
	duckyCodeBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Could not find the provided file in your filesystem or there was an error reading from it.")
		os.Exit(1)
	}
	duckyCode := string(duckyCodeBytes)

	// Look through every line in seperate and convert it
	for _, line := range strings.Split(duckyCode, "\n") {
		fmt.Println(convert(line))
	}
}

// Function to convert the given line
func convert(line string) (string, error) {
	return line, nil
}
