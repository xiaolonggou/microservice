package data

type ArtPiece struct {
	ID             int    `json:"id"`
	Format         string `json:"format"`
	Creator        string `json:"creator"`
	Description    string `json:"-"`
	CreationOn     string `json:"-"`
	LearnedAboutOn string `json:"-"`
}

func GetArtPieceList() []*ArtPiece {
	return artList
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
