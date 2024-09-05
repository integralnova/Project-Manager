package service

type request struct {
	id int
	requester account

}

var request_queue = make(map[int]request)

func WebService() 
	for{
		if request_queue != nil {
			// process request
		} else {
			// wait for request
		}
	}
}
