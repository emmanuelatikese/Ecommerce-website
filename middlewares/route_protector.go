package middlewares

import (
	db_mongo "api/db/mongo"
	jwt_util "api/jwt"
	"api/models"
	response "api/utils"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProtectRoute(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtCookies, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		userCollection := db_mongo.UserCollection
		jwtToken := jwtCookies.Value
		userId, err := jwt_util.Verify(w, r, jwtToken, jwt_util.Access_token_public)
		if err != nil || userId == ""{
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		var filterUser models.User
		priId, err := primitive.ObjectIDFromHex(userId)
		response.ErrorHandler(err, w, 500)
		err = userCollection.FindOne(ctx, bson.M{"_id": priId}, options.FindOne().SetProjection(bson.M{"password": 0})).Decode(&filterUser)
		response.ErrorHandler(err, w, http.StatusUnauthorized)
		ctx = context.WithValue(r.Context(),"User", filterUser)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}