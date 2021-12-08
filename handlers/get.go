package handlers

import (
	"net/http"

	"github.com/xiaolonggou/microservice/v1/data"
)

func (ap *ArtPiece) GetArtPieces(rw http.ResponseWriter, r *http.Request) {
	ap.l.Println("handle http GET ArtPieces request")
	la := data.GetArtPieceList()
	error := la.ToJson(rw)

	if error != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
