package data

import (
	"encoding/json"
	"io"

)

type PrivacySetting struct {
	IsPublic         bool     `json:"is_public"`
	IsInboxOpen      bool     `json:"is_inbox_open"`
	IsTaggingAllowed bool     `json:"is_tagging_allowed"`
	Graylist         []Profile `json:"graylist" gorm:"many2many:profile_graylisted;"`
	Blacklist        []Profile `json:"blacklist" gorm:"many2many:profile_blacklisted;"`
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