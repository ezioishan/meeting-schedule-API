package main

import (
	"fmt"
	"log"
	"net/http"
)
func getMeetingByTime() {
	fmt.Printf("get meetings by time!\n")
}
func getMeetingByParticipant() {
	fmt.Printf("get meetings by participant!\n")
}
func scheduleMeeting() {
	fmt.Printf("Schedule meeting\n")
}
func handleMeetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "//Meetings Endpoint\n")
	if r.Method == "GET" {
		query := r.URL.Query()
		participant := query.Get("participant")

		if participant == "" {
			// start := query.Get("start")
			// end := query.Get("end")
			getMeetingByTime()
		} else {
			getMeetingByParticipant()
		}
		
	} else if r.Method == "POST" {
		scheduleMeeting()
	}

}

func getMeeting(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	fmt.Fprintf(w, "//Meeting Endpoint\n")
	fmt.Fprintf(w, "id is : %s", id)
	// w.WriteHeader(200)
    // w.Write([]byte(id))
}

func handleRequests() {
	http.HandleFunc("/meeting", getMeeting)
	http.HandleFunc("/meetings", handleMeetings)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("hello, world in Go")
	handleRequests()
}
