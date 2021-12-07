package data

import (
	"encoding/json"
	"io"
)

type ArtPiece struct {
	ID             int    `json:"id"`
	Format         string `json:"format"`
	Creator        string `json:"creator"`
	Description    string `json:"-"`
	CreationOn     string `json:"-"`
	LearnedAboutOn string `json:"-"`
}

type ArtPieces []*ArtPiece

func (a *ArtPiece) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func (a *ArtPieces) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(a)
}

func GetArtPieceList() ArtPieces {
	return artList
}

func AddArtPiece(ap *ArtPiece) {
	ap.ID = getNextID()

	artList = append(artList, ap)
}

func getNextID() int {
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
		Creator:        "Knef",
		Description:    "es soll für mich rote Rosen regnen",
		CreationOn:     "09-10-1960",
		LearnedAboutOn: "01-12-2021",
	},
}
