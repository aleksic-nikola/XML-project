package handlers

import(
	"log"
	"net/http"
	"xml/interaction-service/data"
)

type ProfileNotifications struct {
	l *log.Logger
}

func NewProfileNotifications(l *log.Logger) *ProfileNotifications {
	return &ProfileNotifications{l}
}

func(m *ProfileNotifications) GetProfileNotifications(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET Request in profileNotificationHandler")

	lp := data.GetPostNotifications()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}