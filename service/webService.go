package service

import (
	"github.com/Lookm/goProject/project_manager/repl"
	"fmt"
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
	repl()
}