package handlers

import (
	"log"
)

type PrivacySettings struct {
	l *log.Logger
}

func NewPrivacySettings(l *log.Logger) *PrivacySettings {
	return &PrivacySettings{l}
}