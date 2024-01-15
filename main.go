package main

import (
	"fmt"
	"os"
	"pingviewparser/parser"
	"time"
)

func main() {
	var (
		err       error      // error
		filename  string     // filename from args
		parsedtgu [][]string // parsed phone numbers
	)

	// Welcome and usage prompt
	fmt.Println("PingView Template Generator")
	fmt.Println("===================================")
	fmt.Println("")
	fmt.Println("Usage: pingviewparser <filename>")
	fmt.Println("")
	fmt.Println("Example: pingviewparser input.csv")
	fmt.Println("===================================")
	fmt.Println("")

	// check if enough arguments
	if len(os.Args) < 2 {
		panic("Not enough arguments")
	}

	// populate variables
	filename = os.Args[1]

	// start timer for execution time
	start := time.Now()

	// create output directory recursively
	fmt.Println("Generating output folders...")
	err = os.MkdirAll("out/template", 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")

	// parse csv
	fmt.Println("Parsing CSV...")
	parsedtgu, err = parser.ParseCsv(filename)
	if err != nil {
		return
	}
	fmt.Println("Done")

	// cycle through parsed data and write template and query files
	fmt.Println("Writing template file...")
	err = parser.WritePingviewTemplate(parsedtgu)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")

	fmt.Println("All done")
	fmt.Println("Check /out directory for results")

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)

}
