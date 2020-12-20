package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
         )

type Users struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	DateOfBirth    string             `json:"dob,omitempty" bson:"dob,omitempty"`
	PhoneNumber    int                `json:"phn,omitempty" bson:"phn,omitempty"`
	EmailAddress   string             `json:"email,omitempty" bson:"email,omitempty"`
	CreationTimeStamp int          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

var client *mongo.Client

func GetUsersEndpoint(response http.ResposeWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ :=primitive.ObjectIDFromHex(params["id"])
	var users Users
	collection := client.Database("test").Collection("users")
	ctx, _ :=context.WithTimeout(context.Background(),10*time.Second)
	err := collection.FindOne(ctx, Users{ID: id}).Decode(users)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" :"` + err.Error() + `"}`))
		return 
	}
json.NewEncoder(respose).Encode(users)
}

func CreateUsersEndpoint(response http.ResponseWriter,request *http.Request){
	response.Header().Set("content-type","application/json")
	var users Users
	json.NewEncoder(request.Body).Decode(&users)
	collection := client.Database("test").Collection("users")
	ctx, _ :=context.WithTimeout(context.Background(),10*time.Second)
	result, _ :=collection.InsertOne(ctx, users)
	json.NewEncoder(response).Encode(result)
	}

	func main(){
		fmt.Println("Starting the application..")
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dubeymoulshri:9425867630@cluster0.bt0ff.mongodb.net/test"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
		router :=mux.NewRouter()
		router.HandleFunc("/users", CreateUsersEndpoint).Methods("POST")
		router.HandleFunc("/users/{id}", GetUsersEndpoint).Methods("GET")
		http.ListenAndServe(":12345",router)	
	}