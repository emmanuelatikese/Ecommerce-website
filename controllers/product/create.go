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
	response.ErrorHandler(err, w, http.StatusBadRequest)
	if newProduct.Name == "" || newProduct.Price == 0 {
		http.Error(w, "fill the required name and price", http.StatusNotAcceptable)
		return
	}
	if newProduct.Image != ""{
		cld, err := cloudinary.New()
		response.ErrorHandler(err, w, 500)
		image, err := cld.Image(newProduct.Image)
		response.ErrorHandler(err, w, 500)
		imageUrl, err := image.String()
		response.ErrorHandler(err, w, 500)
		newProduct.Image = imageUrl
	}

	insertId, err := productCollection.InsertOne(ctx, newProduct)
	response.ErrorHandler(err, w, 500)
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