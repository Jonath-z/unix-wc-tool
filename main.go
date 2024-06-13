package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
		return false
	}

	return false
}

func getByteNumber(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(fileInfo.Size(), fileName)
}

func getNumberOfLines(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}

	fmt.Println(count, fileName)
}

func countWords(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()
	wordCounter := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCounter++
	}

	fmt.Println(wordCounter, fileName)
}

func ProcessFile(fileName *string, cb func(string)) {
	if *fileName == "" {
		panic("Must specify a file Name")
	}

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	pathsInTheCurrentDir := []string{}
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
		} else {
			pathsInTheCurrentDir = append(pathsInTheCurrentDir, path)
		}

		return err
	})

	if err != nil {
		panic(err.Error())
	}

	if contains(pathsInTheCurrentDir, *fileName) {
		panic("This file is found in the current directory")
	}

	cb(*fileName)
}

func main() {
	byteNumber := flag.String("c", "", "Flag to get the file byte number")
	lineNumber := flag.String("l", "", "Flag to get the file line number")
	wordCounter := flag.String("w", "", "Flag to get the file word counter")
	flag.Parse()

	if *byteNumber != "" {
		ProcessFile(byteNumber, getByteNumber)
		return
	}

	if *lineNumber != "" {
		ProcessFile(lineNumber, getNumberOfLines)
		return
	}

	if *wordCounter != "" {
		ProcessFile(wordCounter, countWords)
	}

}
