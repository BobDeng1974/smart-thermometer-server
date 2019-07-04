package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Record struct {
	ID          *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Temperature float32             `json:"temperature" bson:"temperature"`
	Humidity    float32             `json:"humidity" bson:"humidity"`
	AirPressure float32             `json:"airPressure" bson:"airPressure"`
}

func CreateRecord(temperature float32, humidity float32, airPressure float32) Record {
	return Record{Temperature: temperature, Humidity: humidity, AirPressure: airPressure}
}

func SaveRecord(collection *mongo.Collection, r Record) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, r)
	if err != nil {
		log.Fatal(err)
	}

	return res.InsertedID
}
