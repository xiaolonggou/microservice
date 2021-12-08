package handlers

import (
	"net/http"

	"github.com/xiaolonggou/microservice/v1/data"
)

func (ap *ArtPiece) AddArtPiece(rw http.ResponseWriter, r *http.Request) {
	ap.l.Println("handle http POST ArtPiece request")

	apiece := r.Context().Value(KeyArtPiece{}).(*data.ArtPiece)

	data.AddArtPiece(apiece)
}
