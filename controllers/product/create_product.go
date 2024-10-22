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
	response.ErrorHandler(err, w, http.StatusBadRequest)
	if new_product.Name == "" || new_product.Price == 0 {
		http.Error(w, "fill the required name and price", http.StatusNotAcceptable)
		return
	}
	if new_product.Image != ""{
		cld, err := cloudinary.New()
		response.ErrorHandler(err, w, 500)
		image, err := cld.Image(new_product.Image)
		response.ErrorHandler(err, w, 500)
		imageUrl, err := image.String()
		response.ErrorHandler(err, w, 500)
		new_product.Image = imageUrl
	}

	insertId, err := product_collection.InsertOne(ctx, new_product)
	response.ErrorHandler(err, w, 500)
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