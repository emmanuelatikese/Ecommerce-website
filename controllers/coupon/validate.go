package coupon_controllers

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidateCoupon(w http.ResponseWriter, r *http.Request){
	couponId := mux.Vars(r)["coupon_id"]
	if couponId == ""{
		http.Error(w, "No id provided", http.StatusNotAcceptable)
		return
	}
	ctx, couponCollection :=r.Context(), db_mongo.CouponCollection
	user := response.GetUserFromContext(r)
	priId, err := primitive.ObjectIDFromHex(couponId)
	response.ErrorHandler(err, w, 500)
	var filterCoupon models.Coupons
	err = couponCollection.FindOne(ctx, bson.M{"_id": priId, "userid": user.Id, "isactive": true}).Decode(&filterCoupon)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			http.Error(w, "No coupons available", 404)
			return
		}
		http.Error(w, err.Error(), 500)
		return
	}
	if filterCoupon.Expiration.Before(time.Now()){
		response.JsonResponse("Coupon is Invalid", w, http.StatusNotAcceptable)
		return
	}
	response.JsonResponse("Coupon is valid", w, http.StatusOK)
}