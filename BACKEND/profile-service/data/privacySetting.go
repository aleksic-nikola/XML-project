package data

import (
	"encoding/json"
	"io"

)

type PrivacySetting struct {
	IsPublic         bool     `json:"is_public"`
	IsInboxOpen      bool     `json:"is_inbox_open"`
	IsTaggingAllowed bool     `json:"is_tagging_allowed"`
}

func (ps *PrivacySetting) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ps)
}

func (ps *PrivacySettings) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ps)
}

type PrivacySettings []*PrivacySetting