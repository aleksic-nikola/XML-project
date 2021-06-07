package handlers

import (
	"log"
)

type Images struct {
	l *log.Logger
}

func NewImages(l *log.Logger) *Images {
	return &Images{l}
}


