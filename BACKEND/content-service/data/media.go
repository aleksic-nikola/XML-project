package data

import (
	"encoding/json"
	"io"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	//MediaID uint      `gorm:"primaryKey"`
	PostID     uint      `json:"post_refer"`
	Type       MediaType `json:"type"`
	Path       string    `json:"path"`
	LocationID uint      `json:"location_id"`
	Location   Location  `json:"location" gorm:"foreignKey:LocationID"`
}

func (m *Media) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}

func (m *Medias) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

type Medias []*Media
