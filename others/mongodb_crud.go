package main

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name string
	Age  int
	City string
}
type Users struct {
	Collection *mongo.Collection
	UserInfo   map[string](*User)
}

var database = "mongocrud"
var table = "cruds"
var user = initMongo()

func initMongo() *Users {
	users := &Users{}
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

	fmt.Println("Connected to MongoDB!")

	// 获得在数据库里面User集合的handle
	collection := client.Database(database).Collection(table)
	users.Collection = collection
	users.UserInfo = map[string](*User){}
	return users
}
func insert(name string, age int, city string) {
	insertResult, err := user.Collection.InsertOne(context.TODO(), bson.D{
		{"name", name}, {"age", age}, {"city", city},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
func update(name string, age int, city string) {
	// filter：条件，根据name字段修改其他字段的值
	filter := bson.D{{"name", name}}

	update := bson.D{
		// inc：增加，在原来age基础上增大传入的值
		{"$inc", bson.D{
			{"age", age},
		}},
		// set：赋值，将city字段赋予传入的值
		{"$set", bson.D{
			{"city", city},
		}},
	}
	updateResult, err := user.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
func selectOne(name string, age int, city string) {
	filter := bson.D{{"name", name}, {"age", age}, {"city", city}}
	fmt.Println("filter:", filter)
	var result Trainer
	err := user.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}
func selectLot(name string, age int, city string) {
	findOptions := options.Find()
	filter := bson.D{{"name", name}}
	filter = append(filter, bson.E{Key: "city", Value: city})
	filter = append(filter, bson.E{Key: "age", Value: bson.M{"$gte": age}})

	fmt.Println("filter:", filter)
	// Here's an array in which you can store the decoded documents
	results := make([]User, 0)
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := user.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		fmt.Println("err find")
	}
	var resultsM []bson.M
	if err = cur.All(context.TODO(), &resultsM); err != nil {
		fmt.Println("err All")
	}
	// 遍历的顺序是随机的
	for _, result := range resultsM {
		fmt.Println("********", result)
		for key, value := range result {
			fmt.Println(key, ":", value)
		}
	}
	fmt.Println("resultsM: ", resultsM)
}
func delete(name string) {
	deleteResult, err := user.Collection.DeleteOne(context.TODO(), bson.D{{"name", name}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
func main() {
	insert("test", 1, "11")
	selectLot("test", 1, "11")
	delete("test2")
}
