package product_controllers

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"net/http"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == ""{
		http.Error(w ,"Id not found", 404)
		return
	}
	ctx, product_collection := r.Context(), db_mongo.ProductCollection
	prim_id, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		http.Error(w,err.Error(), 500)
		return
	}
	var filter_product models.Product

	err = product_collection.FindOne(ctx, bson.M{"_id": prim_id}).Decode(&filter_product)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			http.Error(w, "No product found to delete", 404)
			return
		}
		http.Error(w,err.Error(), 500)
		return
	}
	if filter_product.Image != ""{
		cld, err := cloudinary.New()
		if err != nil{
			http.Error(w, err.Error(), 500)
			return
		}
		public_id := response.ExtractPublicID(filter_product.Image)
		_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})
		if err != nil {
			http.Error(w, "Failed to delete image: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	del_result, err := product_collection.DeleteOne(ctx, bson.M{"_id": prim_id})
	if err != nil {
		http.Error(w,err.Error(), 500)
		return
	}

	if del_result.DeletedCount == 0 {
		http.Error(w, "No result found ", 404)
		return
	}

	response.JsonResponse("product deleted successfully", w, 200)
}