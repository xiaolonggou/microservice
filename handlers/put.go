package handlers

import (
	"net/http"

	"github.com/xiaolonggou/microservice/v1/data"
)

// swagger:route PUT /arts arts UpdateArtPiece
// Update an art piece details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse
// Update handles POST requests and updates items from the database
func (ap *ArtPiece) Update(rw http.ResponseWriter, r *http.Request) {

	apiece := r.Context().Value(KeyArtPiece{}).(*data.ArtPiece)

	err := data.UpdateArtPiece(apiece)
	if err == data.ErrorArtPieceNotFound {
		rw.WriteHeader(http.StatusNotFound)
		data.ToJson(&GenericError{Message: "Product not found in database"}, rw)
		return
	} else if err != nil {
		rw.WriteHeader(http.StatusNoContent)
		return
	}
}
