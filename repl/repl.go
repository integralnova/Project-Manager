package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

///IGNORE THIS
///IGNORE THIS
///IGNORE THIS
///IGNORE THIS

func scanNext(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return cleanInput(scanner.Text())
}

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
exit:
	for {
		fmt.Print(">> ")
		input := scanNext(scanner)
		text := input[0]
		switch text {
		case "permits":
			//
		case "new":
			// Create a new permit
		case "search":
		case "exit":
			break exit
		default:
			fmt.Println("HUH?")
		}
	}
	fmt.Println("Exiting...")
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
