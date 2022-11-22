package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	var users []map[string]interface{}

	users = append(users, map[string]interface{}{
		"name":    "dilara",
		"surname": "görüm",
	})
	users = append(users, map[string]interface{}{
		"name":    "samet",
		"surname": "ileri",
	})

	bytes, _ := json.Marshal(users)

	http.HandleFunc("/users", func(w http.ResponseWriter, rq *http.Request) {
		w.Write(bytes)
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln(err)
	}
}

// k6 run --vus 10 --duration 30s script.js

//Metrics:
//vus: Current number of active virtual users
//vus_max: Max possible number of virtual users
//iterations: The aggregate number of times the VUs executed the JS script (the default function).
//iteration_duration: The time it took to complete one full iteration
//data_received:The amount of received data.
//data_sent:The amount of data sent

//HTTP-specific built-in metrics:
//http_reqs: How many total HTTP requests k6 generated.
//http_req_blocked: Time spent blocked (waiting for a free TCP connection slot) before initiating the request.
//http_req_connecting: Time spent establishing TCP connection to the remote host.
//http_req_tls_handshaking: Time spent handshaking TLS session with remote host
//http_req_sending: Time spent sending data to the remote host.
//http_req_waiting: Time spent waiting for response from remote host
//http_req_receiving: Time spent receiving response data from the remote host.
//http_req_duration: Total time for the request. It's equal to http_req_sending + http_req_waiting + http_req_receiving
//http_req_failed: the total number of failed requests/ The rate of failed requests a

// rakyll/hey
//hey -n 300 http://localhost:3000/users