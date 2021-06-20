package dto

import (
	"encoding/json"
	"io"
)

type BlockmuteDTO struct {
	Username string `json:"username"`
	UsernameToBlockMute string `json:"usernametoblockmute"`
}

func (b *BlockmuteDTO) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

func (b *BlockmuteDTO) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}