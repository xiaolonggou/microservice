package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-hclog"
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
	// required: true
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

// ArtPieceList returns a list of art pieces
type ArtPieces []*ArtPiece

func (a *ArtPiece) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func ToJson(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

func UpdateArtPiece(ap *ArtPiece) error {
	id := ap.ID
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

func DeleteArtPiece(id int, l hclog.Logger) error {
	_, i, err := findArtPiece(id)

	if err != nil {
		return err
	}

	l.Debug("delete index:", i)
	l.Debug("array index:", len(artList))

	newList := make([]*ArtPiece, len(artList)-1)
	indexOld := 0
	for indexNew := 0; indexOld <= len(artList)-1; indexOld++ {
		if indexOld != i {
			newList[indexNew] = artList[indexOld]
			indexNew++
		}
	}

	artList = newList
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
	{
		ID:             1,
		Format:         "oil on canvas",
		Creator:        "Vincent Van Gogh",
		Description:    "young fisherman on the beach",
		CreationOn:     "09-10-1860",
		LearnedAboutOn: "25-11-2021",
	},

	{
		ID:             2,
		Format:         "song",
		Creator:        "Hildegard Knef",
		Description:    "es soll für mich rote Rosen regnen",
		CreationOn:     "09-10-1960",
		LearnedAboutOn: "01-12-2021",
	},
}
