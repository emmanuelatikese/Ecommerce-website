package stats_controller

import (
	db_mongo "api/db/mongo"
	response "api/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DailyAnalytics(w http.ResponseWriter, r *http.Request){
	endDate := time.Now()
	startDate := endDate.AddDate(0,0,-7)

	// dateList := response.DateRange(startDate, endDate)
	ctx, orderCollection := r.Context(), db_mongo.OrderCollection
	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{
				{Key: "createdat", Value: bson.D{
					{Key: "$gte", Value: startDate},
					{Key: "$lt", Value: endDate},
				}},
			}},
		},
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "$dateToString", Value: bson.D{
						{Key: "format", Value: "%Y-%m-%d"},
						{Key: "date", Value: "$createdAt"},
					}},
				}},
				{Key: "DailyRevenue", Value: bson.D{{Key: "$sum", Value: "$totalamount"}}},
				{Key: "DailySales", Value: bson.D{{Key: "$sum", Value: 1}}},
			}},
		},
		{
			{Key: "$sort", Value: bson.D{
				{Key: "_id", Value: 1},
			}},
		},
	}

	cur, err := orderCollection.Aggregate(ctx, pipeline)
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	var SevenDayOrders []bson.M
	err = cur.All(ctx, &SevenDayOrders)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}

	response.JsonResponse(SevenDayOrders, w, 200)

}