package data

import (
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/go-playground/validator.v9"
)

// ArtPiece defines the structure for an API product
// swagger:model
type ArtPiece struct {
	// the id for the art piece
	//
	// required: false
	// min: 1
	ID int `json:"id"`
	// the format of the art piece
	//
	// required: true
	Format string `json:"format" validate:"required"`
	// the creator of the art piece
	//
	// required: true
	Creator string `json:"creator" validate:"required"`
	// the last deal of the art piece
	//
	// required: false
	LastSoldat int `json:"price" validate:"gte=0"`
	// the description of the art piece
	//
	// required: false
	Description string `json:"description" validate:"required,description"`
	// the creation time of the art piece
	//
	// required: false
	CreationOn string `json:"-"`
	// learned about time of the art piece
	//
	// required: false
	LearnedAboutOn string `json:"-"`
}

type ArtPieces []*ArtPiece

func (a *ArtPiece) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func (a *ArtPiece) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("description", validateDesc)
	return validate.Struct(a)
}

func validateDesc(fl validator.FieldLevel) bool {
	descStr := fl.Field().String()

	if len(descStr) > 5 {
		return true
	}

	return false
}

func (a *ArtPieces) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(a)
}

func UpdateArtPiece(id int, ap *ArtPiece) error {
	_, pos, err := findArtPiece(id)
	if err != nil {
		return err
	}

	//input ID is replaced by generated ID
	ap.ID = id
	artList[pos] = ap

	return nil
}

var ErrorArtPieceNotFound = fmt.Errorf("art piece not found")

func findArtPiece(id int) (*ArtPiece, int, error) {
	for i, ap := range artList {
		if ap.ID == id {
			return ap, i, nil
		}
	}

	return nil, -1, ErrorArtPieceNotFound
}

func GetArtPieceList() ArtPieces {
	return artList
}

func DeleteArtPiece(id int) error {
	_, i, err := findArtPiece(id)

	if err != nil {
		return err
	}

	if i < len(artList)-1 {
		artList = append(artList[:i], artList[i+1])
	} else if i == len(artList)-1 {
		artList = artList[0 : len(artList)-1]
	}

	return nil
}

func AddArtPiece(ap *ArtPiece) {
	ap.ID = getNextID()

	artList = append(artList, ap)
}

func getNextID() int {

	if len(artList) == 0 {
		return 1
	}

	ap := artList[len(artList)-1]
	return ap.ID + 1
}

var artList = []*ArtPiece{
	&ArtPiece{
		ID:             1,
		Format:         "oil on canvas",
		Creator:        "Vincent Van Gogh",
		Description:    "young fisherman on the beach",
		CreationOn:     "09-10-1860",
		LearnedAboutOn: "25-11-2021",
	},

	&ArtPiece{
		ID:             2,
		Format:         "song",
		Creator:        "Hildegard Knef",
		Description:    "es soll für mich rote Rosen regnen",
		CreationOn:     "09-10-1960",
		LearnedAboutOn: "01-12-2021",
	},
}
