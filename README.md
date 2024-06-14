# Unix wc Tool

## Overview
This project is a Unix/Lunix `wc` written in go from [Coding Challenge](https://codingchallenges.fyi/challenges/challenge-wc/). It processes text files to provide various statistics, including the number of bytes, lines, words, and characters.

## Features

- **Byte count (`-c`)**: Outputs the number of bytes in the file.
- **Line count (`-l`)**: Outputs the number of lines in the file.
- **Word count (`-w`)**: Outputs the number of words in the file.
- **Character count (`-m`)**: Outputs the number of characters in the file.

## Installation

1. **Prerequisites**: Ensure you have Go installed. If not, download and install it from the [official Go website](https://golang.org/dl/).
2. **Clone the repository**:
   ```sh
     git clone https://github.com/Jonath-z/unix-wc-tool
   ```
3. **Navigate to the project directory**:
   ```sh
   cd unix-wc-tool
   ```

## Test

You can test with various flags to get the desired file statistics. Below are some examples:

### Examples

1. **Get the number of bytes in a file**:
   ```sh
   go run main.go -c test.txt
   ```
2. **Get the number of lines in a file**:
   ```sh
   go run main.go -l test.txt
   ```
3. **Get the number of words in a file**:
   ```sh
   go run main.go -w test.txt
   ```
4. **Get the number of characters in a file**:
   ```sh
   go run main.go -m test.txt
   ```


## Code Overview

The code is organized into several functions to handle different tasks:

- `contains(s []string, str string) bool`: Checks if a string is present in a slice of strings.
- `getByteNumber(fileName string) int64`: Returns the size of the file in bytes.
- `getNumberOfLines(fileName string) int`: Returns the number of lines in the file.
- `countWords(fileName string) int`: Returns the number of words in the file.
- `counterCharacter(fileName string) int`: Returns the number of characters in the file.
- `ProcessFile(fileName string, cb func(string))`: Processes the file by checking its existence and executing a callback function.
- `main()`: Parses command-line flags and invokes appropriate functions based on the flags provided.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.
