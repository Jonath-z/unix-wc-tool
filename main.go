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

func getByteNumber(fileName string) int64 {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		panic(err.Error())
	}

	return fileInfo.Size()
}

func getNumberOfLines(fileName string) int {
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

	return count
}

func countWords(fileName string) int {
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

	return wordCounter
}

func counterCharacter(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()
	characterCounter := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		characterCounter++
	}

	return characterCounter
}

func ProcessFile(fileName string, cb func(string)) {
	if fileName == "" {
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

	if contains(pathsInTheCurrentDir, fileName) {
		panic("This file is found in the current directory")
	}

	cb(fileName)
}

func main() {
	byteNumber := flag.String("c", "", "Flag to get the file byte number")
	lineNumber := flag.String("l", "", "Flag to get the file line number")
	wordCounter := flag.String("w", "", "Flag to get the file word counter")
	characterCounter := flag.String("m", "", "Flag to get the file character counter")

	flag.Parse()

	if *byteNumber != "" {
		ProcessFile(*byteNumber, func(s string) {
			fmt.Println(getByteNumber(s), s)
		})
		return
	}

	if *lineNumber != "" {
		ProcessFile(*lineNumber, func(s string) {
			fmt.Println(getNumberOfLines(s), s)
		})
		return
	}

	if *wordCounter != "" {
		ProcessFile(*wordCounter, func(s string) {
			fmt.Println(countWords(s), s)
		})
	}

	if *characterCounter != "" {
		ProcessFile(*characterCounter, func(s string) {
			fmt.Println(counterCharacter(s), s)
		})
		return
	}

	ProcessFile(flag.Arg(0), func(s string) {
		byteNumber := getByteNumber(s)
		lineNumber := getNumberOfLines(s)
		wordNumber := countWords(s)

		fmt.Println(byteNumber, lineNumber, wordNumber, s)
	})

}
