package coupon_controllers

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCoupons(w http.ResponseWriter, r *http.Request){
	user := response.GetUserFromContext(r)
	ctx, couponCollection := r.Context(), db_mongo.CouponCollection
	var filterCoupons models.Coupons
	cursor, err := couponCollection.Find(ctx, bson.M{"isactive": true, "userid": user.Id})
	if err != nil {
		if err == mongo.ErrNoDocuments{
			http.Error(w, "No coupons available", 404)
			return
		}
		http.Error(w, err.Error(), 500)
		return
	}
	err = cursor.All(ctx, &filterCoupons)
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	response.JsonResponse(filterCoupons, w, 200)
}