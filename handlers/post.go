package handlers

import (
	"net/http"

	"github.com/xiaolonggou/microservice/v1/data"
)

// swagger:route POST /arts arts AddArtPiece
// Create a new art piece
//
// responses:
//	200: artPieceResponse
//  422: errorValidation
//  501: errorResponse

func (ap *ArtPiece) Create(rw http.ResponseWriter, r *http.Request) {
	ap.l.Debug("handle http POST ArtPiece request")

	apiece := r.Context().Value(KeyArtPiece{}).(*data.ArtPiece)

	data.AddArtPiece(apiece)
}
