package main

import (
	"fmt"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo/options"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"strconv"
	"context"
	"time"
)
// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/mongo/readpref"
// "go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options" "context" "time"

type Meeting struct {
	Id				string			`json:"id"`
	Title			string			`json:"title"`
	Participants	*Participants	`json:"participants"`
	StartTime 		string			`json:"startTime"`
	EndTime 		string			`json:"endTime"`
}
type Participants struct {
	Name	string		`json:"name"`
	Email	string		`json:"email"`
	Rsvp	string		`json:"rsvp"`
}

var tasks []Meeting
// var client *mongo.Client

// func getMeetingByTime(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("get meetings by time!\n")
// }
// func getMeetingByParticipant(w http.ResponseWriter, r *http.Request, participant string) {
// 	fmt.Printf("get meetings by participant!\n")
// }
// func scheduleMeeting(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Printf("Schedule meeting\n")
// 	w.Header().Add("content-type", "application/json")
// 	var meeting Meeting
// 	json.NewDecoder(r.Body).Decode(&meeting)
// 	collection := client.Database("MeetingDatabse").Collection("Meetings")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	result, _ := collection.InsertOne(ctx, meeting);
// 	json.NewEncoder(w).Encode(result)

// }
func handleMeetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// fmt.Fprintf(w, "/Meetings Endpoint\n")
	if r.Method == "GET" {
		query := r.URL.Query()
		participant := query.Get("participant")

		if participant == "" {
			// start := query.Get("start")
			// end := query.Get("end")
			// getMeetingByTime(w, r)
		} else {
			// getMeetingByParticipant(w, r, participant)
		}
		
	} else if r.Method == "POST" {
		// scheduleMeeting(w, r)
		var client *mongo.Client
		var task Meeting
		 _ = json.NewDecoder(r.Body).Decode(&task)
		task.Id = strconv.Itoa(rand.Intn(100))
		collection := client.Database("restapi").Collection("meetings")
		ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
		if err != nil {
			fmt.Print("Here!")
			log.Fatal(err)
		}
		result, _ := collection.InsertOne(ctx, task)
		// tasks = append(tasks, task)
		json.NewEncoder(w).Encode(result)
	}

}

func getMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	query := r.URL.Query()
	id := query.Get("id")
	fmt.Fprintf(w, "//Meeting Endpoint\n")
	fmt.Fprintf(w, "id is : %s", id)
	for _, item := range tasks {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Meeting{})
}

func main() {
	fmt.Println("Server started...")

	tasks = append(tasks, Meeting{Id: "1", Title: "abc", Participants: &Participants{Name: "ishan", Email: "x@m.com", Rsvp: "yes"}, StartTime:"3pm", EndTime: "4pm"})
	tasks = append(tasks, Meeting{Id: "2", Title: "ezzz", Participants: &Participants{Name: "john", Email: "z@m.com", Rsvp: "yes"}, StartTime:"5pm", EndTime: "6pm"})
	
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
        log.Fatal(err)
    }
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
        log.Fatal(err)
    }
	// get collection as ref
	// collection := client.Database("restapi").Collection("meetings")

	http.HandleFunc("/meeting", getMeeting)
	http.HandleFunc("/meetings", handleMeetings)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
