package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xiaolonggou/microservice/v1/data"
)

// swagger:route DELETE /art/{id} artpieces DeleteArtPiece
// Delete an art piece
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

func (ap *ArtPiece) DeleteArtPiece(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	ap.l.Debug("Handle DELETE art piece", id)

	err := data.DeleteArtPiece(id, ap.l)

	if err == data.ErrorArtPieceNotFound {
		http.Error(rw, "Art not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Art not found", http.StatusInternalServerError)
		return
	}

}
