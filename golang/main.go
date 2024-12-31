package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	countBytesCmd := flag.Bool("c", false, "ccwc -c path/to/file - Returns the number of bytes in the file specified")
	countLinesCmd := flag.Bool("l", false, "ccwc -l path/to/file - Returns the number of lines in the file specified")
	countWordsCmd := flag.Bool("w", false, "ccwc -w path/to/file - Returns the number of words in the file specified")
	countCharactersCmd := flag.Bool("m", false, "ccwc -m path/to/file - Returns the number of characters in the file specified")
	flag.Parse()

	var err error
	var fileName string
	var fileSource *os.File

	if len(os.Args) == 1 || len(os.Args) == 2 && (os.Args[1] == "-c" || os.Args[1] == "-l" || os.Args[1] == "-w" || os.Args[1] == "-m") {
		fileSource = os.Stdin
	} else {
		fileName = os.Args[len(os.Args)-1]
		fileSource, err = os.Open(fileName)
		if err != nil {
			fmt.Printf("error opening the file '%s': %v", fileName, err)
			return
		}
	}

	defer func() {
		closeErr := fileSource.Close()
		if closeErr != nil {
			fmt.Printf("error closing file or stdin: %v\n", closeErr)
		}
	}()

	result, err := getStatsFromReader(fileSource)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *countBytesCmd {
		fmt.Printf("  %d %s\n", result.TotalBytes, fileName)
		return
	}

	if *countCharactersCmd {
		fmt.Printf("  %d %s\n", result.TotalCharacters, fileName)
		return
	}

	if *countLinesCmd {
		fmt.Printf("    %d %s\n", result.TotalLines, fileName)
		return
	}

	if *countWordsCmd {
		fmt.Printf("   %d %s\n", result.TotalWords, fileName)
		return
	}

	fmt.Printf("    %d   %d  %d %s\n", result.TotalLines, result.TotalWords, result.TotalBytes, fileName)
}

type Result struct {
	TotalLines      int64
	TotalBytes      int64
	TotalWords      int
	TotalCharacters int64
}

func getStatsFromReader(r io.Reader) (Result, error) {
	result := Result{
		TotalLines:      0,
		TotalBytes:      0,
		TotalWords:      0,
		TotalCharacters: 0,
	}

	reader := bufio.NewReader(r)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return result, fmt.Errorf("error reading line: %w", err)
		}

		result.TotalLines++
		result.TotalBytes += int64(len(line))
		result.TotalWords += len(strings.Fields(string(line)))
		result.TotalCharacters += int64(utf8.RuneCount(line))

		if errors.Is(err, io.EOF) {
			break
		}
	}

	return result, nil
}
