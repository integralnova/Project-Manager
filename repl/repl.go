package repl

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	Db "github.com/integralnova/Project-Manager/models"
	_ "github.com/mattn/go-sqlite3"
)

func scanNext(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return cleanInput(scanner.Text())
}

type app struct {
	permits *Db.Datatings
}

func Repl() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		permits: &Db.Datatings{
			DB: db,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
exit:
	for {
		fmt.Println("Enter Help to see commands. Enter a command:")
		fmt.Print(">> ")
		input := scanNext(scanner)
		text := input[0]
		switch text {
		case "help":
			printCommands()
		case "permits":
			app.showperimts()
		case "new":
			newpermit(app, input, scanner)
		case "newcompany":
			p := Db.PermitModelPermitCompany{Permit: input[1], CompanyName: input[2]}
			err := app.permits.UpdatePermitCompany(p)
			if err != nil {
				fmt.Println(err)
			}
		case "newdesigner":
			p := Db.PermitModelPermitDesigner{Permit: input[1], Designer: input[2]}
			err := app.permits.UpdatePermitDesigner(p)
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

func spaceadder(b int) string {
	x := new(space)
	normalized := x.normalize(b)
	a := ":"
	for i := normalized; i > 0; i-- {
		a += "."
	}
	return a
}

type space int

func (b *space) normalize(i int) int {
	i = i * i
	i = int(math.Sqrt(float64(i)))
	return i
}

func printCommands() {
	for i := range Commands {
		a := spaceadder(len(i) - (longestcommand + 5))
		fmt.Println(i, a, Commands[i])
	}
}

func (app *app) showperimts() {
	a, _ := app.permits.Getpermits()
	for _, i := range a {
		fmt.Println(i)
	}
}

func newpermit(app app, input []string, scanner *bufio.Scanner) {

	if len(input) < 2 {
		fmt.Println("Please enter a permit number")
		input = append(input, scanNext(scanner)...)
	}

	p := Db.PermitModelPermitID{Permit: input[1]}
	err := app.permits.InsertPermit(p)
	if err != nil {
		fmt.Println(err)
	}
}
