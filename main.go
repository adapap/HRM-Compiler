package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"hrm/compiler"
)

func main() {
	// Enter level number and path to level source code
	// Optionally include -debug flag (for development/testing)
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable compiler debug mode.")
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Printf("Usage: hrm <level> <source path>\n")
		os.Exit(1)
	}
	level, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Printf("Level must be a positive integer.\n")
		os.Exit(1)
	}
	// Handle reading file source
	path := flag.Arg(1)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	source := string(bytes)
	if len(source) == 0 {
		fmt.Printf("No data read from '%s'.\n", path)
		os.Exit(1)
	}
	// Test level by compiling and comparing with expected values
	hrm.TestLevel(level, source, debug)
}
