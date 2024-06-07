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
}

func main() {
	fileName := flag.String("-c", "", "Specify the file")
	flag.Parse()

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

	//if contains(pathsInTheCurrentDir, fileName)

}
