package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xiaolonggou/microservice/v1/data"
)

func (ap *ArtPiece) DeleteartPiece(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	ap.l.Println("Handle DELETE art piece", id)

	err := data.DeleteArtPiece(id)

	if err == data.ErrorArtPieceNotFound {
		http.Error(rw, "Art not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Art not found", http.StatusInternalServerError)
		return
	}

}
