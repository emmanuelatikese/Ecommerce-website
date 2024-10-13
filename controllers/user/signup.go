package user_controllers

import (
	db_mongo "api/db/mongo"
	redis_db "api/db/redis"
	jwt_util "api/jwt"
	"api/models"
	response "api/utils"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	userCollection := db_mongo.UserCollection
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

	err = userCollection.FindOne(ctx, bson.M{"username": new_user_map["username"]}).Decode(&user_exist)
	if err != mongo.ErrNoDocuments{
		http.Error(w, "User already exist", http.StatusNotAcceptable)
		return
	}


	hashed_password, err := bcrypt.GenerateFromPassword([]byte(new_user_map["password"]), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if new_user_map["role"] == ""{
		new_user_map["role"] = "customer"
	}
	new_user = models.User{
		Username: new_user_map["username"],
		Email: new_user_map["email"],
		Password: hashed_password,
		CartItem: []models.Cart{},
		Role: new_user_map["role"],
	}

	insert_id, err := userCollection.InsertOne(ctx, new_user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
		log.Println("Here")
	priId, ok := insert_id.InsertedID.(primitive.ObjectID)
	if !ok{
		http.Error(w, "can't convert id", 500)
	}
	strId := priId.Hex()

	access_token, refresh_token, err := jwt_util.GenerateToken(strId)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	redis_db.SetRfToken(strId, refresh_token, ctx)
	// log.Println("there")
	jwt_util.SetCookie(access_token, refresh_token, w)

	res := map[string]interface{}{
		"_id": insert_id.InsertedID,
		"username": new_user.Username,
		"email": new_user.Email,
		"cart_item": []models.Cart{},
		"roles": new_user.Role,
	}
	response.JsonResponse(res, w, 201)
}