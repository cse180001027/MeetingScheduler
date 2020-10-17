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


func scheduleMeeting(writeresponse http.ResponseWriter, request *http.Request){
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



	for participant := range (mparticipants) {
		context, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		filter, error := participantcollection.Find(ctx, bson.M{"meetingid":})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
	}

	//update and do things appropriately
	responseBody ,error := http.Post("/meetings", "json",bytes.NewBuffer(meetinginstance)
	fmt.Fprintf(writeresponse, responseBody)

}

func getMeeting(writeresponse http.ResponseWriter,request *http.Request){

	defer request.Body.Close()

	urldata,error := ioutil.ReadAll(request.Body)
	if error != nil {
		panic(error)
	}

	var	meetingid = request.URL.Path[2:]

	var meeting Meeting

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


	filter := bson.M{"meetingid": meetingid }
	context, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = meetingscollection.FindOne(ctx, filter).Decode(&meeting)
	if err != nil {
		log.Fatal(err)
	}

	meet, err := json.Marshal(meeting)
	if err != nil {
		fmt.Println(err)
		return
	}
	//get meeting from database having id obtained from url
	fmt.Fprint(writeresponse,meet)
}




func listMeeting(writeresponse http.ResponseWriter,request *http.Request){

	defer request.Body.Close()

	request.ParseForm()
	starttime := request.Form.Get("start")
	endtime   := request.Form.Get("end")



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


	context, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter, err := collection.Find(ctx, bson.D{})
	if err != nil { log.Fatal(err) }
	defer filter.Close(ctx)
	for filter.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		// do something with result....
	}
	if err := filter.Err(); err != nil {
		log.Fatal(err)
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
