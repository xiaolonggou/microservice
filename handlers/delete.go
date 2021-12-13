package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xiaolonggou/microservice/v1/data"
)

// swagger:route DELETE /art/{id} arts DeleteArtPiece
// Delete an art piece
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse
// Delete handles DELETE requests and removes items from the database
func (ap *ArtPiece) DeleteArtPiece(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	ap.l.Debug("Handle DELETE art piece", id)

	err := data.DeleteArtPiece(id, ap.l)

	if err == data.ErrorArtPieceNotFound {
		http.Error(rw, "Art not found", http.StatusNotFound)
		data.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		http.Error(rw, "Art not found", http.StatusInternalServerError)
		data.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

}
