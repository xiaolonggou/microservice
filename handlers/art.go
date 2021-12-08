package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"xiaolong.com/v2/data"
)

type ArtPiece struct {
	l *log.Logger
}

func NewArtPiece(l *log.Logger) *ArtPiece {
	return &ArtPiece{l}
}

func (ap *ArtPiece) UpdateArtPiece(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	artpiceId, errStrconv := strconv.Atoi(vars["id"])

	if errStrconv != nil {
		http.Error(rw, "Unable to convert id to integer", http.StatusBadRequest)
		return
	}

	ap.l.Println("handle http PUT ArtPiece request")

	apiece := &data.ArtPiece{}

	err := apiece.FromJson(r.Body)

	if err != nil {
		ap.l.Printf("error %#v", err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateArtPiece(artpiceId, apiece)
	if err == data.ErrorArtPieceNotFound {
		http.Error(rw, "Art piece not found.", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Art piece not found.", http.StatusInternalServerError)
		return
	}
}

func (ap *ArtPiece) GetArtPieces(rw http.ResponseWriter, r *http.Request) {
	ap.l.Println("handle http GET ArtPieces request")
	la := data.GetArtPieceList()
	error := la.ToJson(rw)

	if error != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (ap *ArtPiece) AddArtPiece(rw http.ResponseWriter, r *http.Request) {
	ap.l.Println("handle http POST ArtPiece request")

	apiece := &data.ArtPiece{}

	err := apiece.FromJson(r.Body)

	if err != nil {
		ap.l.Printf("error %#v", err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddArtPiece(apiece)
}
