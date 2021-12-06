package handlers

import (
	"log"
	"net/http"

	"xiaolong.com/v2/data"
)

type ArtPiece struct {
	l *log.Logger
}

func NewArtPiece(l *log.Logger) *ArtPiece {
	return &ArtPiece{l}
}

func (ap *ArtPiece) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	la := data.GetArtPieceList()
	error := la.ToJson(rw)

	if error != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
