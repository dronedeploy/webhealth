package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
    "flag"
)

var lastCheckin = time.Now()
var currentlyHealthy = true
var heartbeat = 10
var grace = 1

func healthcheck(w http.ResponseWriter, r *http.Request) {
    if currentlyHealthy {
		io.WriteString(w, "ok")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		io.WriteString(w, "not ok")
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	lastCheckin = time.Now()
	io.WriteString(w, "pong")
}

func updateStatus() {
	for {
	    if time.Now().Before(lastCheckin.Add(time.Duration(heartbeat * grace) * time.Second)) {
		    currentlyHealthy = true
			fmt.Println("healthy ", time.Now())
		} else {
		    currentlyHealthy = false
			fmt.Println("not healthy ", time.Now())
		}
		time.Sleep(time.Duration(heartbeat) * time.Second)
	}
}

//func main() {
//	fmt.Println("starting: ", time.Now())
//	go updateStatus()
//	http.HandleFunc("/health", healthcheck)
//	http.HandleFunc("/ping", ping)
//	http.ListenAndServe(":8000", nil)
//}

func init() {
    flag.IntVar(&heartbeat, "heartbeat", 10, "heartbeat interval in seconds")
    flag.IntVar(&grace, "grace", 10, "number of intervals that can be missed before considered unhealthy")
    flag.Parse()
}
