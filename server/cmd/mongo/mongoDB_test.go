package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 设置账户密码
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "123456",
	}).ApplyURI("mongodb://127.0.0.1:27017"))
	// clientOptions := options.Client().ApplyURI("mongodb://admin:123456@127.0.0.1:27017")
	// client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return client.Database("lucar")
}

// func TestMongoInsert(t *testing.T) {
// 	c := context.Background()
// 	// mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	mdb := InitMongo()
// 	col := mdb.Collection("account")
// 	res, err := col.InsertMany(c, []interface{}{
// 		bson.M{
// 			"open_id": "678",
// 		},
// 		bson.M{
// 			"open_id": "456",
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v", res)
// }

// func TestFindRows(t *testing.T) {
// 	c := context.Background()
// 	mdb := InitMongo()
// 	col := mdb.Collection("account")
// 	res := col.FindOne(c, bson.M{
// 		"open_id": "456",
// 	})
// 	var row struct {
// 		ID     primitive.ObjectID `bson:"_id"`
// 		OpenID string             `bson:"open_id"`
// 	}
// 	err := res.Decode(&row)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", row)
// }

func TestFindRowsByCur(t *testing.T) {
	c := context.Background()
	mdb := InitMongo()
	col := mdb.Collection("account")
	cur, err := col.Find(c, bson.M{})
	if err != nil {
		panic(err)
	}
	for cur.Next(c) {
		var row struct {
			ID     primitive.ObjectID `bson:"_id"`
			OpenID string             `bson:"open_id"`
		}
		err = cur.Decode(&row)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", row)
	}
}
