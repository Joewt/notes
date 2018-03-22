package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occured on openint the inputfile\n" + 
			"Does the file exist?\n" +
			"Have you got acces to it ?\n")
		return 
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("The input was: %s", inputString)
	} 
}
