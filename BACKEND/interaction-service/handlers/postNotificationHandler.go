package handlers

import(
	"log"
	"net/http"
	"xml/interaction-service/data"
)

type PostNotifications struct {
	l *log.Logger
}

func NewPostNotifications(l *log.Logger) *PostNotifications {
	return &PostNotifications{l}
}

func(m *PostNotifications) GetPostNotifications(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET Request in postNotificationHandler")

	lp := data.GetPostNotifications()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}