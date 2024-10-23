package product_controllers

import (
	db_mongo "api/db/mongo"
	redis_db "api/db/redis"
	"api/models"
	response "api/utils"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateFeatured(w http.ResponseWriter, r *http.Request){
	Id := mux.Vars(r)["id"]
	if Id == ""{
		http.Error(w, "No category was passed", 404)
		return
	}
	ctx, productCollection := r.Context(), db_mongo.ProductCollection
	strId, err  := primitive.ObjectIDFromHex(Id)
	response.ErrorHandler(err, w, 500)
	var filterProduct models.Product
	err = productCollection.FindOne(ctx, bson.M{"_id": strId}).Decode(&filterProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			http.Error(w, "Product don't exist", 404)
			return
		}
		http.Error(w, err.Error(), 404)
		return
	}
	filterProduct.IsFeatured = !filterProduct.IsFeatured
	_, err = productCollection.UpdateOne(ctx, bson.M{"_id": strId}, bson.M{"$set":bson.M{"isfeatured": filterProduct.IsFeatured}})
	response.ErrorHandler(err, w, http.StatusNotAcceptable)
	var all_featured []bson.M
	cur, err := productCollection.Find(ctx, bson.M{"isfeatured": true})
	response.ErrorHandler(err, w, http.StatusNotAcceptable)
	err = cur.All(ctx, &all_featured);
	response.ErrorHandler(err, w, 500)
	redis_db.RedisCli.Set(ctx, "featured_product", all_featured, time.Hour * 24 * 7)
	response.JsonResponse(filterProduct, w, 200)
}