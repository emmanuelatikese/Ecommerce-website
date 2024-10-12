package controllers

import (
	db_mongo "api/db"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var new_user models.User
	new_user_map := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&new_user_map)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserCollection := db_mongo.UserCollection
	ctx := r.Context()

	if len(new_user_map["password"]) < 6{
		http.Error(w, "password must be 6 or more characters", http.StatusNotAcceptable)
		return
	}

	if new_user_map["username"] == "" || new_user_map["password"] != new_user_map["confirmed_password"] {
		http.Error(w, "Invalid username or password mismatch", http.StatusNotAcceptable)
		return
	}
	var user_exist models.User
	err = UserCollection.FindOne(ctx, bson.M{"username": new_user_map["username"]}).Decode(&user_exist)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			http.Error(w, "User already exist", http.StatusNotAcceptable)
			return
		}
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(new_user_map["password"]), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	new_user = models.User{
		Username: new_user_map["username"],
		Email: new_user_map["email"],
		Password: hashed_password,
		CartItem: []models.Cart{},
		Role: "customer",
	}

	insert_id, err := UserCollection.InsertOne(ctx, new_user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	res := map[string]interface{}{
		"_id": insert_id,
		"username": new_user.Username,
		"email": new_user.Email,
		"cart_item": []models.Cart{},
		"roles": "customer",
	}
	response.JsonResponse(res, w, 201)
}