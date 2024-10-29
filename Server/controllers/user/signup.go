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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	newUserMmap := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&newUserMmap)
	isErr := response.ErrorHandler(err, w, http.StatusBadRequest)
	if isErr{
		return
	}
	userCollection := db_mongo.UserCollection
	ctx := r.Context()
	if len(newUserMmap["password"]) < 6{
		http.Error(w, "password must be 6 or more characters", http.StatusNotAcceptable)
		return
	}
	if newUserMmap["username"] == "" || newUserMmap["password"] != newUserMmap["confirmed_password"] {
		http.Error(w, "Invalid username or password mismatch", http.StatusNotAcceptable)
		return
	}
	var user_exist models.User
	err = userCollection.FindOne(ctx, bson.M{"username": newUserMmap["username"]}).Decode(&user_exist)
	if err != mongo.ErrNoDocuments{
		http.Error(w, "User already exist", http.StatusNotAcceptable)
		return
	}
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(newUserMmap["password"]), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if newUserMmap["role"] == ""{
		newUserMmap["role"] = "customer"
	}
	newUser = models.User{
		Username: newUserMmap["username"],
		Email: newUserMmap["email"],
		Password: hashed_password,
		CartItem: []models.Cart{},
		Role: newUserMmap["role"],
	}
	insert_id, err := userCollection.InsertOne(ctx, newUser)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	priId, ok := insert_id.InsertedID.(primitive.ObjectID)
	if !ok{
		http.Error(w, "can't convert id", 500)
	}
	strId := priId.Hex()
	access_token, refresh_token, err := jwt_util.GenerateToken(strId)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	redis_db.SetRfToken(strId, refresh_token, ctx)
	jwt_util.SetCookie(access_token, refresh_token, w)
	res := map[string]interface{}{
		"_id": insert_id.InsertedID,
		"username": newUser.Username,
		"email": newUser.Email,
		"cartitem": []models.Cart{},
		"role": newUser.Role,
	}
	response.JsonResponse(res, w, 201)
}