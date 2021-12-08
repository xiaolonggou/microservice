package data

import "testing"

func TestCheckValidation(t *testing.T) {
	a := &ArtPiece{}
	a.Format = "1"
	a.Creator = "Max Mustermann"
	a.LastSoldat = 10000
	a.Description = "Unkown Picture"
	err := a.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
