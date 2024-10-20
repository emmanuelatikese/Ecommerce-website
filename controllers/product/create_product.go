package product_controllers

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"
	"github.com/cloudinary/cloudinary-go/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(w http.ResponseWriter, r *http.Request){
	var new_product models.Product
	ctx, product_collection := r.Context(), db_mongo.ProductCollection

	err := json.NewDecoder(r.Body).Decode(&new_product)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if new_product.Name == "" || new_product.Price == 0 {
		http.Error(w, "fill the required name and price", http.StatusNotAcceptable)
		return
	}
	if new_product.Image != ""{
		cld, err := cloudinary.New()
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		image, err := cld.Image(new_product.Image)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		imageUrl, err := image.String()
		if err != nil {
				http.Error(w, err.Error(), 500)
				return
		}
		new_product.Image = imageUrl
	}

	insertId, err := product_collection.InsertOne(ctx, new_product)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pri_Id, ok := insertId.InsertedID.(primitive.ObjectID)
	if!ok{
		http.Error(w, "Unable to convert Id", 500)
		return
	}
	strId := pri_Id.Hex()

	res := &bson.M{
		"_id": strId,
		"name": new_product.Name,
		"price": new_product.Price,
		"image": new_product.Image,
		"description": new_product.Description,
		"isfeatured": new_product.IsFeatured,
		"createdat": new_product.CreatedAt,
		"updatedat": new_product.UpdatedAt,
		"category": new_product.Category,
	}

	response.JsonResponse(res, w, 201)
}