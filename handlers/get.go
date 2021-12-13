package handlers

import (
	"net/http"

	"github.com/xiaolonggou/microservice/v1/data"
)

// swagger:route GET /arts arts listArtPieces
// Returns a list of art pieces
// responses:
//	200: artPieceResponse

// GetArtPieces returns the art pieces from the data store
func (ap *ArtPiece) GetArtPieces(rw http.ResponseWriter, r *http.Request) {
	ap.l.Debug("handle http GET ArtPieces request")
	rw.Header().Add("Content-Type", "application/json")

	la := data.GetArtPieceList()

	err := data.ToJson(la, rw)
	if err != nil {
		// we should never be here but log the error just incase
		ap.l.Error("Unable to serializing product", "error", err)
	}
}
