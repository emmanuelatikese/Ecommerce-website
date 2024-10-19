package product_controllers

import (
	db_mongo "api/db/mongo"
	response "api/utils"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
// NB: test this using postman after adding more product

func GetRecommendation(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	product_collection := db_mongo.ProductCollection

pipeline := mongo.Pipeline{
    {
        {Key: "$sample", Value: bson.D{{Key: "size", Value: 3}}},
    },
    {
        {Key: "$project", Value: bson.D{
            {Key: "_id", Value: 1},
            {Key: "name", Value: 1},
            {Key: "description", Value: 1},
            {Key: "image", Value: 1},
        }},
    },
}

	cursor, err := product_collection.Aggregate(ctx, pipeline)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	var recommendation []bson.M
	if err = cursor.All(ctx, &recommendation); err != nil{
		http.Error(w, err.Error(), 500)
		return
	}

	response.JsonResponse(recommendation, w, 200)
}