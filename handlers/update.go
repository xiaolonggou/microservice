package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xiaolonggou/microservice/v1/data"
)

func (ap *ArtPiece) UpdateArtPiece(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	artpiceId, errStrconv := strconv.Atoi(vars["id"])

	if errStrconv != nil {
		http.Error(rw, "Unable to convert id to integer", http.StatusBadRequest)
		return
	}

	ap.l.Debug("handle http PUT ArtPiece request")

	apiece := r.Context().Value(KeyArtPiece{}).(*data.ArtPiece)

	err := data.UpdateArtPiece(artpiceId, apiece)
	if err == data.ErrorArtPieceNotFound {
		http.Error(rw, "Art piece not found.", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Art piece not found.", http.StatusInternalServerError)
		return
	}
}
