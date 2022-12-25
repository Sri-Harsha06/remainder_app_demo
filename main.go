package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Event struct {
	Id        string `json:"id,omitempty" bson:"id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Event     string `json:"event,omitempty" bson:"event,omitempty"`
	Date      string `json:"date,omitempty" bson:"date,omitempty"`
	Time      string `json:"time,omitempty" bson:"time,omitempty"`
	CreatedAt string `json:"createdat,omitempty" bson:"createdat,omitempty"`
	UpdatedAt string `json:"updatedat,omitempty" bson:"updatedat,omitempty"`
	CreatedBy string `json:"createdby,omitempty" bson:"createdby,omitempty"`
	UpdatedBy string `json:"updatedby,omitempty" bson:"updatedby,omitempty"`
}


func main() {
	// fmt.Println(time.Now().AddDate(0, 0, 1).Format("02-01-2006"))
	// fmt.Println(time.Now().Format("15:04"))
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://harsha:harsha@cluster0.xnvctix.mongodb.net/?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/addevent", AddEvent).Methods("POST")
	router.HandleFunc("/allevents", getEvents).Methods("GET")
	router.HandleFunc("/event/{id}", readEventById).Methods("GET")
	router.HandleFunc("/event2/{name}", readEventByName).Methods("GET")
	router.HandleFunc("/event3/{event}", readEventByEvent).Methods("GET")
	router.HandleFunc("/event4/{date}", readEventByDate).Methods("GET")
	router.HandleFunc("/update", updateEvent).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteEvent).Methods("GET")
	router.HandleFunc("/eventstom", findtmrevents).Methods("GET")
	http.ListenAndServe(":12345", router)
}

func AddEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var event Event
	// fmt.Println(request.Body)
	_ = json.NewDecoder(request.Body).Decode(&event)
	// fmt.Println(event)
	collection := client.Database("events").Collection("event")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, event)
	json.NewEncoder(response).Encode(result)
}

func readEventById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	// fmt.Printf("%T", id)
	var event Event
	collection := client.Database("events").Collection("event")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Event{Id: id}).Decode(&event)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(event)
}

func getEvents(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []Event
	collection := client.Database("events").Collection("event")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var event Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func readEventByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []Event
	params := mux.Vars(request)
	name := params["name"]
	// fmt.Printf("%T", id)
	collection := client.Database("events").Collection("event")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, Event{Name: name})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var event Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func readEventByEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []Event
	params := mux.Vars(request)
	event := params["event"]
	collection := client.Database("events").Collection("event")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, Event{Event: event})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var event Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func readEventByDate(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []Event
	params := mux.Vars(request)
	date := params["date"]
	collection := client.Database("events").Collection("event")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, Event{Date: date})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var event Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func updateEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	collection := client.Database("events").Collection("event")
	var event Event
	_ = json.NewDecoder(request.Body).Decode(&event)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{Key: "id", Value: event.Id}}
	replacement := Event(event)
	result, err := collection.ReplaceOne(ctx, filter, replacement)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(response).Encode(result)
}

func deleteEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	// fmt.Printf("%T", id)
	collection := client.Database("events").Collection("event")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		panic(err)
	}
	json.NewEncoder(response).Encode(result)
}

func findtmrevents(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []Event
	collection := client.Database("events").Collection("event")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	dt := time.Now().AddDate(0, 0, 1).Format("02-01-2006")
	fmt.Printf("%T", dt)
	fmt.Println(dt)
	cursor, err := collection.Find(ctx, Event{Date: dt})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var event Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}
