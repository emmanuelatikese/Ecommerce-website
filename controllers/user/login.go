package user_controllers

import (
	db_mongo "api/db/mongo"
	redis_db "api/db/redis"
	jwt_util "api/jwt"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request){
    user_logins := make(map[string]string)
    ctx := r.Context()
    user_collection := db_mongo.UserCollection
    err := json.NewDecoder(r.Body).Decode(&user_logins)
    if err != nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var filter_user models.User
    err = user_collection.FindOne(ctx, bson.M{"username": user_logins["username"]}).Decode(&filter_user)
    if err != nil{
        if err == mongo.ErrNoDocuments{
            http.Error(w, "Account don't exist", http.StatusNotAcceptable)
            return
        }
        http.Error(w, err.Error(), http.StatusNotAcceptable)
        return
    }

    err = bcrypt.CompareHashAndPassword(filter_user.Password, []byte(user_logins["password"]))
    if err != nil {
        http.Error(w, "password invalid", http.StatusNotAcceptable)
        return
    }
    access_token, refresh_token, err := jwt_util.GenerateToken(filter_user.Id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	redis_db.SetRfToken(filter_user.Id.String(), refresh_token, ctx)
	jwt_util.SetCookie(access_token, refresh_token, w)

    response.JsonResponse(filter_user, w, 201)
}