// Package classification of Product API
//
// Documentation for Art Piece API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//  swagger:meta
package handlers

import "github.com/xiaolonggou/microservice/v1/data"

// A list of art pieces
// swagger:response artPiecesResponse
type artPiecesResponseWrapper struct {
	// All art pieces in the system
	// in: body
	Body []data.ArtPiece
}
