package controllers

import (
	"net/http"

	"github.com/unsign3d/smart-thermometer-server/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func AddRecord(w http.ResponseWriter, r *http.Request) {
	client := utils.MongoClient()
	collection := client.Database("sensor-data").Collection("sensor-data")
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	w.WriteHeader(http.StatusCreated)
}
