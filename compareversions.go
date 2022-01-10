package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {

	// CLI usage info
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Println("compareversions 1.3.4 1.2.4")
	}
	flag.Parse()

	// Check if have the correct number of arguments
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	// Read command-line parameters
	ver1 := args[0]
	ver2 := args[1]

	// Run helper function
	result, err := compareVersions(ver1, ver2)
	if err != nil {
		fmt.Println("Error occurred while running compareVersions function, check the arguments:", err)
		os.Exit(1)
	}

	// Return result to stdout
	fmt.Fprint(os.Stdout, result)
}

// Compares two version strings
func compareVersions(version1, version2 string) (int, error) {

	// Split version strings into slices
	versionQue1 := strings.Split(version1, ".")
	versionQue2 := strings.Split(version2, ".")

	for len(versionQue1) > 0 || len(versionQue2) > 0 {
		// Cases if one version have more levels
		if len(versionQue1) == 0 {
			return -1, nil
		}
		if len(versionQue2) == 0 {
			return 1, nil
		}

		// Read the first element of the slice, return an error if can not convert to int
		number1, err := strconv.Atoi(versionQue1[0])
		if err != nil {
			return 0, err
		}
		number2, err := strconv.Atoi(versionQue2[0])
		if err != nil {
			return 0, err
		}

		// Compare the current level of revision, return a result if not equal
		if number1 > number2 {
			return 1, nil
		}
		if number1 < number2 {
			return -1, nil
		}

		// Continue checking the next level of revision
		versionQue1 = versionQue1[1:]
		versionQue2 = versionQue2[1:]
	}

	// Return 0 if no difference in versions
	return 0, nil
}