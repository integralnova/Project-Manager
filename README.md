<h1> Project Tracker </h1>

<p>Simple Web App to track Permits</p>

<h2> Dependencies</h2>'
<p> GCC </p>
<p> GOOSE
go install golang.org/x/tools/gopls@latest

go install github.com/pressly/goose/v3/cmd/goose@latest
goose goose -dir=assets/migrations sqlite3 app.db up 

</p>