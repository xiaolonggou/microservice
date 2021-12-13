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
// swagger:response artPieceResponse
type artPieceResponseWrapper struct {
	// All art pieces in the system
	// in: body
	Body []data.ArtPiece
}

// swagger:parameters DeleteArtPiece
type artIDParamsWrapper struct {
	// The id of the art piece for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}
