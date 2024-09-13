package repl

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	sqlite "github.com/integralnova/Project-Manager/dbManager"
	_ "github.com/mattn/go-sqlite3"
)

func scanNext(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return cleanInput(scanner.Text())
}

type app struct {
	permits *sqlite.PermitModel
}

func Repl() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		permits: &sqlite.PermitModel{
			DB: db,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
exit:
	for {
		fmt.Print(">> ")
		input := scanNext(scanner)
		text := input[0]
		switch text {
		case "permits":
			a, _ := app.permits.Getpermits()
			for _, i := range a {
				fmt.Println(i)
			}
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
