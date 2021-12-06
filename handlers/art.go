package handlers

import (
	"encoding/json"
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
	d, err := json.Marshal(la)

	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
