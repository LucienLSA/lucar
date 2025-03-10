package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IDField defines the field name for mongo document id.
const IDField = "_id"

// ObjID defines the object id field
type ObjectID struct {
	ID primitive.ObjectID `bson:"_id"`
}

// Set returns a $set update document.
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert returns a $setOnInsert update document.
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
