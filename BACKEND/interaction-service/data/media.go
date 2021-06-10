package data

import (
	"encoding/json"
	"io"
)

type Media struct {
	ID   int       `json:"id"`
	Path string    `json:"path"`
	Type MediaType `json:"type"`
}

func (m *Media) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(e)
}

// collection
type Medias []*Media

func (m *Medias) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func GetMedias() Medias {
	return mediaList
}

var mediaList = []*Media{
	{
		ID:   1,
		Path: "myphath/myimage",
		Type: image,
	},
	{
		ID:   2,
		Path: "myphath/myvideo",
		Type: video,
	},
}
