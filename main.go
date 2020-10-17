package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"bytes"
	"io/ioutil"
	"time"
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)


type Meeting struct{
	Id string
	Title string
	Participants map[string]bool
	Start time.Time
	End time.Time
	Creationtime time.Time
}



type Participant struct{
	Name string
	Email string
	RSVP map[string]bool
}




func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func scheduleMeeting(w http.ResponseWriter, request *http.Request){
	requestdecode := json.NewDecoder(request.Body)
	// Declare a new Meeting struct.

	var meetinginstance Meeting
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := requestdecode.Decode(&meetinginstance)
	if err != nil {
		panic(err)
	}
	var mparticipants = meetinginstance.Participants


	context,cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter, error := participantcollection.Find(ctx, bson.M{})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}









	filterCursor, error := participantcollection.Find(ctx, bson.M{"RSVP": })
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M{}
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)



	context , cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reseult, error := collection.InsertOne(context, bson.M{"": "pi", "value": 3.14159})
	id := res.InsertedID










	json.Unmarshal( requestdecode []byte, &meetinginstance)


	error := json.NewDecoder(r.Body).Decode(&p)
	if (error != nil){
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	//update and do things appropriately
	responseBody ,error := http.Post("/meetings", "json",bytes.NewBuffer(meetinginstance)
	fmt.Fprintf(w, "Person: %+v", )

}

func getMeeting(writeresponse http.ResponseWriter,request *http.Request){

	defer request.Body.Close()

	urldata, error := ioutil.ReadAll(request.Body)
	if error != nil {
		panic(error)
	}

	//get meeting from database having id obtained from url

	fmt.Fprint(writeresponse,meeting)
}




func listMeeting(writeresponse http.ResponseWriter,request *http.Request){

	defer request.Body.Close()

	urldata, error := ioutil.ReadAll(request.Body)
	if error != nil {
		panic(error)
	}

	//get start and end times and return
	//the meetings having start time or end time
	//in the interval
	//meeting from database having id obtained from url

	fmt.Fprint(writeresponse,meetingarray)
}


func participantMeetings(writeresponse http.ResponseWriter,request *http.Request){
	defer request.Body.Close()
	urldata, error := ioutil.ReadAll(request.Body)
	if error != nil {
		panic(error)
	}

	//get emailid and return
	//all the meetings
	//from participant map RSVP in  database having starttime>=Now()


	fmt.Fprint(writeresponse,meetingarray)
}


func main() {

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, error := mongo.Connect(context, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if error = client.Disconnect(context); error != nil {
			panic(error)
		}
	}()

	context, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	error = client.Ping(context, readpref.Primary())
	meetingscollection := client.Database("ScheduleMeet").Collection("meetings")
	participantcollection := client.Database("ScheduleMeet").Collection("participants")

	context, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, error := participantcollection.Insert(ctx, bson.M{"name": "John",
		"email": "geop.123@gmail.com"})
	id := res.InsertedID




	//to access values of url in function request.url.Path[start:end]
	mux:= http.NewServeMux()
	mux.HandleFunc("/meetings",scheduleMeeting)
	mux.HandleFunc("/meetings/<id here>",getMeeting)
	mux.HandleFunc("/meetings?start=<start time here>&end=<end time here>",listMeeting)
	mux.HandleFunc("/meetings?participant=<email id>",participantMeetings)

	error := http.ListenAndServe(":4000", mux)
	log.Fatal(error)

	if(error != nil){
		log.Fatalln(error)
	}
}
