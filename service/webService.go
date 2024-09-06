package service

import (
	"fmt"

	"github.com/integralnova/Project-Manager/repl"
)

type request struct {
	id        int
	requester string //user

}

var request_queue = make(map[int]request)

func WebService() {
	start_repl()

}

func start_repl() {
	fmt.Println("Starting web service...")
	repl.Repl()
}
