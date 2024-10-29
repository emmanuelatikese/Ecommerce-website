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
    userLogins := make(map[string]string)
    ctx := r.Context()
    userCollection := db_mongo.UserCollection
    err := json.NewDecoder(r.Body).Decode(&userLogins)
    isErr := response.ErrorHandler(err, w, http.StatusBadRequest)
	if isErr{
		return
	}
    var filterUser models.User
    err = userCollection.FindOne(ctx, bson.M{"username": userLogins["username"]}).Decode(&filterUser)
    if err != nil{
        if err == mongo.ErrNoDocuments{
            http.Error(w, "Account don't exist", http.StatusNotAcceptable)
            return
        }
        http.Error(w, err.Error(), http.StatusNotAcceptable)
        return
    }
    err = bcrypt.CompareHashAndPassword(filterUser.Password, []byte(userLogins["password"]))
    if err != nil {
        http.Error(w, "password invalid", http.StatusNotAcceptable)
        return
    }
    accessToken, refreshToken, err := jwt_util.GenerateToken(filterUser.Id.Hex())
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
    redis_db.SetRfToken(filterUser.Id.Hex(), refreshToken, ctx)
	jwt_util.SetCookie(accessToken, refreshToken, w)
    res := &map[string]interface{}{
        "_id": filterUser.Id,
        "username": filterUser.Username,
        "email": filterUser.Email,
        "cartitem": filterUser.CartItem,
        "role":filterUser.Role,
    }
    response.JsonResponse(res, w, 201)
}