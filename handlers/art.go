package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	} else if r.Method == http.MethodPut {
		//expect id in th uri
		reg := regexp.MustCompile(`/([0-9]+)`)
		group := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(group) != 1 {
			ap.l.Println("not one id")
			http.Error(rw, "invalid URI", http.StatusBadRequest)
		}

		if len(group[0]) != 2 {
			ap.l.Println("not two capture groups:", group[0])
			http.Error(rw, "invalid URI", http.StatusBadRequest)
		}

		idString := group[0][1]

		id, err := strconv.Atoi(idString)

		if err != nil {
			ap.l.Println("string cannot be converted to integer:", idString)
			http.Error(rw, "invalid URI", http.StatusBadRequest)
		}
		ap.updateArtPiece(id, rw, r)
		return
	} else {
		//catch all other cases
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ap *ArtPiece) updateArtPiece(artpiceId int, rw http.ResponseWriter, r *http.Request) {
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
