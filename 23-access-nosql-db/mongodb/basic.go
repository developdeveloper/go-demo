package main

import (
	"23-access-nosql-db/session"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ping() {
	client := session.GetMongoClient()
	defer client.Disconnect(nil)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
}

func insert() {
	client := session.GetMongoClient()
	defer client.Disconnect(nil)

	persons := client.Database("test").Collection("persons")
	persons.InsertOne(nil, bson.M{
		"username":    "zhangsan",
		"nickname":    "张三",
		"age":         20,
		"create_time": time.Now(),
	})
}

func findOne() {
	client := session.GetMongoClient()
	defer client.Disconnect(nil)

	persons := client.Database("test").Collection("persons")
	var result bson.M
	//persons.FindOne(nil, bson.D{}).Decode(&result)
	persons.FindOne(nil, bson.D{{"username", "zhangsan"}}).Decode(&result)
	fmt.Println(result)
}

func find() {
	client := session.GetMongoClient()
	defer client.Disconnect(nil)

	persons := client.Database("test").Collection("persons")
	cur, _ := persons.Find(nil, bson.D{})

	for cur.Next(nil) {
		var result bson.M
		cur.Decode(&result)
		fmt.Println(result)
	}

	// cur.Err()
}

func main() {
	ping()
	//insert()
	findOne()
	find()
}
