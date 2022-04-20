package api

import (
	"context"
	"dev11/internal/model"
	"dev11/internal/storage"
	"fmt"
	"net"
	"net/http"
	"time"
)

// Store Интерфейс для работы с базой данных
type Store interface {
	CreateEvent(event *model.Event) error
	UpdateEvent(event *model.Event) error
	DeleteEvent(event *model.Event)
	EventsForDay(userId int, date time.Time) ([]model.Event, error)
	EventsForWeek(userId int, date time.Time) ([]model.Event, error)
	EventsForMonth(userId int, date time.Time) ([]model.Event, error)
}

type Server struct {
	storage Store
	server  *http.Server
}

func NewServer(storage *storage.DB, port string) *Server {
	return &Server{
		storage: storage,
		server: &http.Server{
			Addr: net.JoinHostPort("", port),
		},
	}
}
func (s *Server) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", s.createEvent)
	mux.HandleFunc("/update_event", s.updateEvent)
	mux.HandleFunc("/delete_event", s.deleteEvent)
	mux.HandleFunc("/events_for_day", s.eventsForDay)
	mux.HandleFunc("/events_for_week", s.eventsForWeek)
	mux.HandleFunc("/events_for_month", s.eventsForMonth)
	handler := Logger(mux)
	return handler
}
func (s *Server) RunServer() error {
	s.server.Handler = s.InitRoutes()
	return s.server.ListenAndServe()
}
func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("shutting down server")
	return s.server.Shutdown(ctx)
}
