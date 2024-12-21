package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
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
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("unable to open file '%s' due to: %v", fileName, err)
	}
	defer func() { _ = f.Close() }()

	buf := make([]byte, 1024)
	var totalBytes int64
	reader := bufio.NewReader(f)

	for {
		n, err := reader.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return fmt.Errorf("error occurred trying to scane the file '%s': %v", fileName, err)
		}

		totalBytes += int64(n)
	}

	// Print the output as specified
	fmt.Printf("%d %s\n", totalBytes, fileName)

	return nil
}
