package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string for local host
const connectionStringLocal = "mongodb://localhost:27017"

const connectionString = "connection String"

//Database name
const DBName = "GoTest"

//collection name
const collName = "toDoList"

//object\Instance
var collection *mongo.Collection

//start connection with mongo db
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect with mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	//check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongoDB")

	collection = client.Database(DBName).Collection(collName)

	fmt.Println("collectecion instance created")
}

func getAllTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask
	json.NewEncoder(w).Encode(payload)
}

//Create task routes
func CreateTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contex-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.toDoList
	_ = json.NewDecoder(r.Body).Decode(&task)

	fmt.Println(task, r.body)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)

}