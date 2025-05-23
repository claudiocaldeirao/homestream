package entity

type OmdbMetadata struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
	Response   string `json:"Response"`
	Error      string `json:"Error"`
}

func NewDefaultOmdbMetadata(title string) *OmdbMetadata {
	return &OmdbMetadata{
		Title:      title,
		Year:       "N/A",
		Rated:      "N/A",
		Released:   "N/A",
		Runtime:    "N/A",
		Genre:      "N/A",
		Director:   "N/A",
		Plot:       "Metadata not found.",
		Language:   "N/A",
		Country:    "N/A",
		Awards:     "N/A",
		Poster:     "",
		Metascore:  "N/A",
		ImdbRating: "N/A",
		ImdbID:     "",
		Response:   "False",
		Error:      "Movie not found in OMDb",
	}
}
