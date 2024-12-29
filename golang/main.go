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

	fileName := os.Args[len(os.Args)-1]

	if *countBytesCmd {
		err := countBytesInFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if *countLinesCmd {
		err := countLinesInFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if *countWordsCmd {
		err := countWordsInFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if *countCharactersCmd {
		err := countCharacterInFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func countLinesInFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err // TODO: refactor me
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	totalLines := 0

	for scanner.Scan() {
		if errors.Is(scanner.Err(), io.EOF) {
			break
		}

		if scanner.Err() != nil {
			return err //TODO
		}

		totalLines++
	}

	fmt.Printf("%d %s\n", totalLines, fileName)

	return nil
}

func countBytesInFile(fileName string) error {
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

func countWordsInFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err // TODO: refactor me
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	totalWords := 0

	for scanner.Scan() {
		if errors.Is(scanner.Err(), io.EOF) {
			break
		}

		if scanner.Err() != nil {
			return err //TODO
		}

		totalWords += len(strings.Fields(scanner.Text()))
	}

	fmt.Printf("%d %s\n", totalWords, fileName)

	return nil
}

// FIXME: off by 4 bytes when compared to wc
func countCharacterInFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err // TODO: refactor me
	}
	defer func() { _ = f.Close() }()

	totalCharacter := 0

	buf := make([]byte, 1024)
	reader := bufio.NewReader(f)

	for {
		n, err := reader.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err //TODO
		}

		totalCharacter += utf8.RuneCount(buf[:n])
	}

	fmt.Printf("%d %s\n", totalCharacter, fileName)

	return nil
}
