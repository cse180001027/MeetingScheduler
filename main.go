package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"bytes"
	"io/ioutil"
	"strings"
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
	Starttime time.Time
	Endtime time.Time
	Creationtime time.Time
}



type Participant struct{
	Name string
	Email string
	RSVP map[string]bool
}


func connecttodb() (){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, error := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if error = client.Disconnect(context); error != nil {
			panic(error)
		}
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	error = client.Ping(ctx, readpref.Primary())

	if error != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	meetingscollection := client.Database("ScheduleMeet").Collection("meetings")
	participantcollection := client.Database("ScheduleMeet").Collection("participants")

	return meetingscollection,participantcollection
}

func sendemailgetrsvp(var emailid string) (status){
	status := "response: \"yes\" No Maybe "
	return status
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
	var mparticipants := meetinginstance.Participants
	var mstart := meetinginstance.Starttime
	var mend := meetinginstance.Endtime
	meetingscollection ,participantscollection := connecttodb()

	var result []struct{
		Email string `bson:"Email"`
		Id string `bson:"Id"`
	}

	for participant := range mparticipants{
		meetinginstance.Participants[participant] = true
	}

	for participant := range (mparticipants){
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := participantscollection.Find(ctx,bson.D{"Name":participant} ,bson.M{"$or": []bson.M{
			bson.M{"RSVP.status":"Yes"},
			bson.M{"RSVP.status":"May be"}}}).Select(bson.M{"Email": 1,"Id" : 1}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		for _, emailids := range result {
			var previousmeet Meeting
			err := json.Unmarshal(getMeeting(w, httppost)[]byte, &previousmeet)
			if err != nil {
				fmt.Println(err)
				return
			}
			if ((previousmeet.Starttime.After(mstart) ||
				previousmeet.Starttime.Equal(mstart)) &&
				(previousmeet.Starttime.Before(mend) ||
					previousmeet.Starttime.Equal(mend)) ||
				(previousmeet.Endtime.After(mstart) ||
					previousmeet.Endtime.Equal(mstart)) &&
					(previousmeet.Endtime.Before(mend) ||
						previousmeet.Endtime.Equal(mend))
		){
			var httppost = "/meetings/" + string(emailids.Id)
			httppost = "\""+httppost+"\""
			status := sendemailgetrsvp(emailids.Email)
			if (strings.ToLower(status) == "yes" || strings.ToLower(status) == "maybe") {
				err := participantscollection.update(
				{
				},
				{
					"$pull": {
					"Name": participant, "RSVP": {
						"ID":  emailids.Id
					}
				}
				},
				{multi:true}
				)
				if (err != nil) {
					panic(err)
				}
			} else {
				meetinginstance.Participants[participant] = false


			}


		} else {



			}
			}

		}
		defer filter.Close(context)
		for filter.Next(context){
			var result bson.M
			err := filter.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			sendemailgetrsvp(result.Retrieve())

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


	neetingscollection , participantcollection := connecttodb()


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
	start := request.Form.Get("start")
	end   := request.Form.Get("end")

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

	var meetingsarray []bson.D

	context, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter , err := meetingscollection.Find(ctx, bson.D{"Starttime": bson.D{"$gte":start,"$lte":end}}, "Endtime": bson.D{"$gte":start,"$lte":end}})
	if err != nil { log.Fatal(err) }


	if err = filter.All(context, &meetingsarray); err != nil {
	panic(err)
	}

	if err = filter.Err(); err != nil {
		log.Fatal(err)
	}
	//get start and end times and return
	//the meetings having start time or end time
	//in the interval
	//meeting from database having id obtained from url
	fmt.Fprint(writeresponse,meetingsarray)
}


func participantMeetings(writeresponse http.ResponseWriter,request *http.Request){
	defer request.Body.Close()

	request.ParseForm()

	emailid := request.Form.Get("participant")

	meetingscollection,participantcollection := connecttodb()

	var meetingsarray []bson.D

	context, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter , err := participantcollection.Find(ctx, bson.D{"Email": emailid })
	if err != nil { log.Fatal(err) }

	if err = filter.All(context , &meetingsarray); err != nil {
		panic(err)
	}
	//get emailid and return
	//all the meetings
	//from participant map RSVP in  database having starttime>=Now()

	fmt.Fprint(writeresponse,meetingsarray)
}


func main() {


	meetingscollection,participantcollection := connecttodb()

	context,cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, error := participantcollection.Insert(ctx, bson.M{"name": "John",
		"email": "geop.123@gmail.com"})
	id := result.InsertedID


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
