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
	if r.Method == http.MethodGet {
		ap.getArtPieces(rw, r)
		return
	} else if r.Method == http.MethodPost {
		ap.addArtPieces(rw, r)
		return
	} else {
		//catch all other cases
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (ap *ArtPiece) getArtPieces(rw http.ResponseWriter, r *http.Request) {
	ap.l.Println("handle http GET ArtPieces request")
	la := data.GetArtPieceList()
	error := la.ToJson(rw)

	if error != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (ap *ArtPiece) addArtPieces(rw http.ResponseWriter, r *http.Request) {
	ap.l.Println("handle http POST ArtPiece request")

	apiece := &data.ArtPiece{}

	err := apiece.FromJson(r.Body)

	if err != nil {
		ap.l.Printf("error %#v", err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddArtPiece(apiece)
}
