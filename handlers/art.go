package handlers

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/xiaolonggou/microservice/v1/data"
)

type ArtPiece struct {
	l hclog.Logger
	v *data.Validation
}

func NewArtPiece(l hclog.Logger, v *data.Validation) *ArtPiece {
	return &ArtPiece{l, v}
}

type KeyArtPiece struct{}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
