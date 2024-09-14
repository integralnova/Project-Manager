package repl

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	models "github.com/integralnova/Project-Manager/internal"
	_ "github.com/mattn/go-sqlite3"
)

func scanNext(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return cleanInput(scanner.Text())
}

type app struct {
	permits *models.APPDB
}

func Repl() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		permits: &models.APPDB{
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
			fmt.Println(a)
			for _, i := range a {
				fmt.Println(i)
			}
		case "new":
			p := models.PermitModelPermitID{Permit: input[1]}
			err := app.permits.NewPermit(p)
			if err != nil {
				fmt.Println(err)
			}
		case "newcompany":
			p := models.PermitModelPermitCompany{Permit: input[1], CompanyName: input[2]}
			err := app.permits.UpdatePermitCompany(p)
			if err != nil {
				fmt.Println(err)
			}
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
