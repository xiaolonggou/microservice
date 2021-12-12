package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidation(t *testing.T) {
	a := &ArtPiece{}
	a.Format = "1"
	a.Creator = "Max Mustermann"
	a.LastSoldat = 10000
	a.Description = "ture"

	v := NewValidation()
	err := v.Validate(a)
	assert.Len(t, err, 1)
}
