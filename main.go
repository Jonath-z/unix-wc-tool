package main

import (
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

type Flags struct {
	byteNumber string
	lineNumber string
}

func main() {
	flags := Flags{
		byteNumber: "c",
		lineNumber: "l",
	}

	fileName := flag.String(flags.byteNumber, "", "Specify the file")
	//fileName = flag.String(flags.lineNumber, "", "Specify the file")

	flag.Parse()

	fmt.Println(flag.Args())
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

	fileInfo, err := os.Stat(*fileName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(fileInfo.Size(), *fileName)
}
