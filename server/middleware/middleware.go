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
	"../models"
)

// DB connection string for local host
const connectionStringLocal = "mongodb://localhost:27017"

//const connectionString = 

//Database name
const DBName = "GoTest"

//collection name
const collName = "ToDoList"

//object\Instance
var collection *mongo.Collection

//start connection with mongo db
func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionStringLocal)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(DBName).Collection(collName)

	fmt.Println("Collection instance created!")
}

// get all the task from the DB
func GetAllTask(w http.ResponseWriter, r *http.Request){
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

	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)

	fmt.Println(task, r.Body)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)

}

// update task route
func TaskComplete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//undo the complte task route
func UndoTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//delete task in route
func DeleteTask( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.newEncoder(w).encode("Task not found")
}

//delete all tasks route
func DeleteAllTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlenconded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
}

//Get all the task from the DB.
func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M 
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)

		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	
	cur.Close(context.Background())
	return results
}

//insert one task in the db
func insertOneTask( task models.ToDoList){
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single Record", insertResult.InsertedID)
}

func taskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

//undo method from task, uptsate the status of the task to false
func undoTask(task string){
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}
// delete task from the db by ID
func deleteOneTask(task string){
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all tht task from the db
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(),bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Documetn", d.DeletedCount)
	return d.DeletedCount
}
