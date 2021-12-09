// Package classification of Product API
//
// Documentation for Art Piece API
//
//  Schemes: http
//  Basepath:/
//  Version: 1.0.0
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//  swagger:meta

package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/xiaolonggou/microservice/v1/data"
)

// A list of art pieces returns in the reponse
// swagger: response artPiecesResponse
type artPiecesResponseWrapper struct {
	// All art pieces in the system
	// in: body
	Body []data.ArtPiece
}

type ArtPiece struct {
	l *log.Logger
}

func NewArtPiece(l *log.Logger) *ArtPiece {
	return &ArtPiece{l}
}

type KeyArtPiece struct{}

func (ap ArtPiece) MiddlewareArtPieceValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			apiece := &data.ArtPiece{}

			err := apiece.FromJson(r.Body)
			if err != nil {
				ap.l.Printf("[Error] %#v", err)
				http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
				return
			}

			//validate the art piece
			err = apiece.Validate()
			if err != nil {
				ap.l.Printf("[Error] %#v", err)
				http.Error(rw, fmt.Sprintf("Unable to validate input json: %s", err), http.StatusBadRequest)
				ap.l.Println(apiece)
				return
			}

			ctx := context.WithValue(r.Context(), KeyArtPiece{}, apiece)
			req := r.WithContext(ctx)
			next.ServeHTTP(rw, req)
		})

}
