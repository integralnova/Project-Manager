package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanNext(scanner bufio.Scanner) []string {
	scanner.Scan()
	return cleanInput(scanner.Text())
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		input := scanNext(scanner)
		if input[0] == "exit" {
			break
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
