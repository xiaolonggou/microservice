package handlers

import (
	"context"
	"net/http"

	"github.com/xiaolonggou/microservice/v1/data"
)

func (ap *ArtPiece) MiddlewareArtPieceValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			apiece := &data.ArtPiece{}

			err := apiece.FromJson(r.Body)
			if err != nil {
				ap.l.Error("[Error] %#v", err)
				http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
				return
			}

			//validate the art piece
			errs := ap.v.Validate(apiece)
			if len(errs) != 0 {
				ap.l.Error("Validating product", "error", errs)

				// return the validation messages as an array
				rw.WriteHeader(http.StatusUnprocessableEntity)
				data.ToJson(&ValidationError{Messages: errs.Errors()}, rw)
				return
			}

			ctx := context.WithValue(r.Context(), KeyArtPiece{}, apiece)
			req := r.WithContext(ctx)
			next.ServeHTTP(rw, req)
		})

}
