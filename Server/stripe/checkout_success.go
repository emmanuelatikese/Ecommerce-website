package stripe_section

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"

	"github.com/stripe/stripe-go/v80/checkout/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckoutSuccess(w http.ResponseWriter, r *http.Request){
	type newSession struct {
		Id string `json:"sessionId"`
	}
	user := response.GetUserFromContext(r)
	var structSession newSession
	err := json.NewDecoder(r.Body).Decode(&structSession)
	isErr := response.ErrorHandler(err, w, http.StatusBadRequest)
	if isErr{
			return
	}
	sessionId := structSession.Id
	if sessionId == ""{
		http.Error(w, "No session Id provided", 404)
		return
	}
	ctx, couponCollection := r.Context(), db_mongo.CouponCollection
	userSession, err := session.Get(sessionId, nil)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
			return
	}
	if userSession.PaymentStatus == "paid"{
		userCoupon := userSession.Metadata["couponCode"]
		if userCoupon != ""{
			var newCouponUpdate models.Coupons
			err := couponCollection.FindOneAndUpdate(ctx, bson.M{"code": userCoupon},
				bson.M{"isactive": false}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&newCouponUpdate)
			isErr = response.ErrorHandler(err, w, 500)
			if isErr{
				return
			}
		}
	}

	var allProduct []models.ProductQty
	err = json.Unmarshal([]byte(userSession.Metadata["products"]), &allProduct)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	var newOrder models.Order
	orderCollection := db_mongo.OrderCollection
	newOrder = models.Order{
		User: user,
		Products: allProduct,
		TotalAmount: float64(userSession.AmountTotal) /100,
		StripeSessionId: sessionId,
	}
	insertId, err := orderCollection.InsertOne(ctx, newOrder)
	isErr = response.ErrorHandler(err, w, 505)
	if isErr{
		return
	}
	orderId, _ := insertId.InsertedID.(*primitive.ObjectID)
	
	res := map[string]interface{}{
		"success": true,
		"message": "payment successfully done",
		"orderId": orderId.Hex(),
	}
	response.JsonResponse(res, w, 201)
}