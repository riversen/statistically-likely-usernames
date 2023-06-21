package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	// Command-line flags for input and output files, and the desired username format
	firstNameFile = flag.String("first", "", "File with first names (This is a required field)")
	lastNameFile  = flag.String("last", "", "File with last names (This is a required field)")
	outputFile    = flag.String("out", "", "Output file for usernames (This is a required field)")
	format        = flag.String("format", "", "Format for usernames (This is a required field)")
)

func main() {
	// Parse the command-line flags
	flag.Parse()

	// Check that all required flags have been specified
	if *firstNameFile == "" || *lastNameFile == "" || *outputFile == "" || *format == "" {
		fmt.Println("You must specify first name file with the -first flag, last name file with the -last flag, output file with the -out flag and format with the -format flag.")
		return
	}

	// Open the input files
	fnameFile, err := os.Open(*firstNameFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fnameFile.Close()

	lnameFile, err := os.Open(*lastNameFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lnameFile.Close()

	// Prepare to scan the input files line by line
	firstNameScanner := bufio.NewScanner(fnameFile)
	lastNameScanner := bufio.NewScanner(lnameFile)

	// Create the output file
	usernames, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer usernames.Close()

	// Read each line (name) from the input files and generate usernames
	for firstNameScanner.Scan() {
		fname := strings.ToLower(firstNameScanner.Text())
		for lastNameScanner.Scan() {
			lname := strings.ToLower(lastNameScanner.Text())
			var username string

			// Generate usernames according to the specified format
			switch *format {
			case "jsmith":
				username = fmt.Sprintf("%s%s\n", string(fname[0]), lname)
			case "johnsmith":
				username = fmt.Sprintf("%s%s\n", fname, lname)
			case "john.smith":
				username = fmt.Sprintf("%s.%s\n", fname, lname)
			case "jjs":
				username = fmt.Sprintf("%s%s\n", string(fname[0]), string(lname[0]))
			case "john":
				username = fmt.Sprintf("%s\n", fname)
			case "smith":
				username = fmt.Sprintf("%s\n", lname)
			case "jjsmith":
				username = fmt.Sprintf("%s%s%s\n", string(fname[0]), string(lname[0]), lname)
			case "smithjj":
				username = fmt.Sprintf("%s%s%s\n", lname, string(fname[0]), string(lname[0]))
			case "johnjs":
				username = fmt.Sprintf("%s%s%s\n", fname, string(fname[0]), string(lname[0]))
			case "smithj":
				username = fmt.Sprintf("%s%s\n", lname, string(fname[0]))
			case "johns":
				username = fmt.Sprintf("%s%s\n", fname, string(lname[0]))
			case "jsmith2":
				username = fmt.Sprintf("%s%s2\n", string(fname[0]), lname)
			case "smithj2":
				username = fmt.Sprintf("%s%s2\n", lname, string(fname[0]))
			default:
				fmt.Printf("Unsupported format: %s\n", *format)
				return
			}

			// Write the generated username to the output file
			_, err := usernames.WriteString(username)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		// Rewind the last name file to the beginning
		lnameFile.Seek(0, 0)
		lastNameScanner = bufio.NewScanner(lnameFile)
	}

	// Check for errors that may have occurred while scanning the input files
	if firstNameScanner.Err() != nil {
		fmt.Println(firstNameScanner.Err())
		return
	}

	if lastNameScanner.Err() != nil {
		fmt.Println(lastNameScanner.Err())
		return
	}
}
