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
	var newProduct models.Product
	ctx, productCollection := r.Context(), db_mongo.ProductCollection
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	isErr := response.ErrorHandler(err, w, http.StatusBadRequest)
	if isErr{
		return
	}
	if newProduct.Name == "" || newProduct.Price == 0 {
		http.Error(w, "fill the required name and price", http.StatusNotAcceptable)
		return
	}
	if newProduct.Image != ""{
		cld, err := cloudinary.New()
		isErr = response.ErrorHandler(err, w, 500)
		if isErr{
			return
		}
		image, err := cld.Image(newProduct.Image)
		isErr = response.ErrorHandler(err, w, 500)
		if isErr{
			return
		}
		imageUrl, err := image.String()
		isErr = response.ErrorHandler(err, w, 500)
		if isErr{
			return
		}
		newProduct.Image = imageUrl
	}

	insertId, err := productCollection.InsertOne(ctx, newProduct)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
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
		"name": newProduct.Name,
		"price": newProduct.Price,
		"image": newProduct.Image,
		"description": newProduct.Description,
		"isfeatured": newProduct.IsFeatured,
		"createdat": newProduct.CreatedAt,
		"updatedat": newProduct.UpdatedAt,
		"category": newProduct.Category,
	}

	response.JsonResponse(res, w, 201)
}