package stats_controller

import (
	db_mongo "api/db/mongo"
	response "api/utils"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAnalytics(w http.ResponseWriter, r *http.Request){
	ctx, userCollection := r.Context(), db_mongo.UserCollection
	productCollection, orderCollection := db_mongo.ProductCollection, db_mongo.OrderCollection
	totalUsers, err := userCollection.CountDocuments(ctx, bson.M{})
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}

	totalProduct, err := productCollection.CountDocuments(ctx, bson.M{})
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}

	pipeline := mongo.Pipeline{
		{
			{
				Key: "$group", Value:  bson.D{
					{Key: "_id", Value: nil},
					{Key: "totalRevenue", Value: bson.D{{Key: "$sum", Value: "$totalamount"}}},
					{Key: "totalSales", Value: bson.D{{Key:"$sum", Value: 1}}},
				},
			},
		},
	}
	cur, err := orderCollection.Aggregate(ctx, pipeline)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	var filterOrder []bson.M
	err = cur.All(ctx, &filterOrder)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	log.Println(filterOrder)
	response.JsonResponse(map[string]interface{}{
		"totalUsers": totalUsers,
		"totalProduct": totalProduct,
	}, w, 200)
}