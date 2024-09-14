package repl

var Commands = map[string]string{
	"help":       "Displays this help message",
	"permits":    "Displays all permits",
	"insert":     "Creates a new a new record. Usage: insert <permitnumber> ",
	"company":    "updates the permit company name. Usage: company <permitnumber> <companyname>",
	"designer":   "updates the permit designer. Date of record creation will be date started. Usage: designer <permitnumber> <designername>",
	"startdate":  "updates the permit start date. Usage: startdate <permitnumber>",
	"finishdate": "updates the permit finish date. Usage: finishdate <permitnumber>",
	"exit":       "Exits the program",
	"note":       "date must be in the format of YYYY-MM-DD",
}

var longestcommand = func() int {
	max := 0
	for i := range Commands {
		if len(i) > max {
			max = len(i)
		}
	}
	return max
}()
