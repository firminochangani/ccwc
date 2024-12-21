package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	countCmd := flag.Bool("c", false, "ccwc -c path/to/file - Returns the number of bytes in the file specified")
	flag.Parse()

	fileName := os.Args[len(os.Args)-1]

	if *countCmd {
		err := countBytes(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func countBytes(fileName string) error {
	return nil
}
