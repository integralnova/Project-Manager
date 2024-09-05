package service

import (
	"fmt"
	"github.com/integralnova/Project-Manager/repl"
)

type request struct {
	id int
	requester string //user

}

var request_queue = make(map[int]request)

func WebService(){
	start_repl()
	for{
		if request_queue != nil {
			// process request
		} else {
			// wait for request
		}
	}
}



func start_repl() {
	fmt.Println("Starting web service...")
	repl.Repl()
}