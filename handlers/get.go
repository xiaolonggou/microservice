package handlers

import (
	"net/http"

	"github.com/xiaolonggou/microservice/v1/data"
)

// swagger:route Get /art pieces listArtPieces
// Returns a list of art pieces
// responses:
//	200: artPieceResponse

// GetArtPieces returns the art pieces from the data store
func (ap *ArtPiece) GetArtPieces(rw http.ResponseWriter, r *http.Request) {
	ap.l.Println("handle http GET ArtPieces request")
	la := data.GetArtPieceList()
	error := la.ToJson(rw)

	if error != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
