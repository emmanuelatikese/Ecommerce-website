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
	ctx, productCollection := r.Context(), db_mongo.ProductCollection
	primId, err := primitive.ObjectIDFromHex(id)
	response.ErrorHandler(err, w, 500)
	var filterProduct models.Product
	err = productCollection.FindOne(ctx, bson.M{"_id": primId}).Decode(&filterProduct)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			http.Error(w, "No product found to delete", 404)
			return
		}
		http.Error(w,err.Error(), 500)
		return
	}
	if filterProduct.Image != ""{
		cld, err := cloudinary.New()
		response.ErrorHandler(err, w, 500)
		public_id := response.ExtractPublicID(filterProduct.Image)
		_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})
		response.ErrorHandler(err, w, 500)
	}
	del_result, err := productCollection.DeleteOne(ctx, bson.M{"_id": primId})
	response.ErrorHandler(err, w, 500)
	if del_result.DeletedCount == 0 {
		http.Error(w, "No result found ", 404)
		return
	}
	response.JsonResponse("product deleted successfully", w, 200)
}