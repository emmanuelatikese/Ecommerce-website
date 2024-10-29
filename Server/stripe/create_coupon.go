package stripe_section

import (
	"api/models"
	response "api/utils"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCoupon(Id primitive.ObjectID, ctx context.Context, coupon *mongo.Collection, w http.ResponseWriter){
    var newCoupon models.Coupons
    newCode := "GIFT"+response.GenerateCouponCode()
    newCoupon = models.Coupons{
        Code: newCode,
        Discount: 20,
        IsActive: true,
        UserId: Id,
    }
    _, err := coupon.InsertOne(ctx, newCoupon)
    response.ErrorHandler(err, w, 500)
}