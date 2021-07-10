package handlers

import "log"

type NotificationSettings struct {
	l *log.Logger
}

func NewNotificationSettings(l *log.Logger) *NotificationSettings {
	return &NotificationSettings{l}
}