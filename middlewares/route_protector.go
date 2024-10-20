package middlewares

import (
	db_mongo "api/db/mongo"
	jwt_util "api/jwt"
	"api/models"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProtectRoute(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt_cookies, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		user_collection := db_mongo.UserCollection
		jwt_token := jwt_cookies.Value
		userId, err := jwt_util.Verify(w, r, jwt_token, jwt_util.Access_token_public)
		if err != nil || userId == ""{
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		var filterUser models.User
		pri_id, err := primitive.ObjectIDFromHex(userId)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = user_collection.FindOne(ctx, bson.M{"_id": pri_id}, options.FindOne().SetProjection(bson.M{"password": 0})).Decode(&filterUser)
		if err != nil{
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx = context.WithValue(r.Context(),"User", filterUser)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}