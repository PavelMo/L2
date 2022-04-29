package api

import (
	"dev11/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	if methodError(w, r.Method, "POST") {
		return
	}

	event, err := s.DecodeJSON(r)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occurred while decoding input data: %v", err), http.StatusBadRequest, true)
		return
	}
	if err = s.storage.CreateEvent(event); err != nil {
		s.Response(w, fmt.Sprintf("error occured: %v", err), http.StatusServiceUnavailable, true)
		return
	}
	s.Response(w, "event created successfully", http.StatusCreated, false)
}
func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	if methodError(w, r.Method, "POST") {
		return
	}
	event, err := s.DecodeJSON(r)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occurred while decoding input data: %v", err), http.StatusBadRequest, true)
		return
	}
	err = s.storage.UpdateEvent(event)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occured while passing data to service for updating: %v", err), http.StatusServiceUnavailable, true)
		return
	}
	s.Response(w, "event updated successfully", http.StatusCreated, false)
}
func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if methodError(w, r.Method, "POST") {
		return
	}
	event, err := s.DecodeJSON(r)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occurred while decoding input data: %v", err), http.StatusBadRequest, true)
		return
	}
	s.storage.DeleteEvent(event)
	s.Response(w, "event deleted successfully", http.StatusCreated, false)
}
func (s *Server) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if methodError(w, r.Method, "GET") {
		return
	}
	userId, date, err := getParams(r.URL)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occured while getting params: %v", err), http.StatusBadRequest, true)
		return
	}
	events, err := s.storage.EventsForDay(userId, date)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occured while getting events for day: %v", err), http.StatusServiceUnavailable, true)
		return
	}
	s.ResponseWithData(w, events)

}
func (s *Server) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if methodError(w, r.Method, "GET") {
		return
	}
	userId, date, err := getParams(r.URL)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occured while getting params: %v", err), http.StatusBadRequest, true)
	}
	events, err := s.storage.EventsForWeek(userId, date)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occured while getting events for week: %v", err), http.StatusServiceUnavailable, true)
		return
	}
	s.ResponseWithData(w, events)
}
func (s *Server) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if methodError(w, r.Method, "GET") {
		return
	}
	userId, date, err := getParams(r.URL)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occured while getting params: %v", err), http.StatusBadRequest, true)
	}
	events, err := s.storage.EventsForMonth(userId, date)
	if err != nil {
		s.Response(w, fmt.Sprintf("error occured while getting events for month: %v", err), http.StatusServiceUnavailable, true)
		return
	}
	s.ResponseWithData(w, events)
}
func (s *Server) ResponseWithData(w http.ResponseWriter, events []model.Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, _ := json.MarshalIndent(events, " ", "")
	_, err := w.Write(res)
	if err != nil {
		s.Response(w, err.Error(), http.StatusInternalServerError, true)
	}
}
func (s *Server) Response(w http.ResponseWriter, message string, status int, isErr bool) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if isErr {
		data := NewErr(message, status)
		resp, _ := json.MarshalIndent(data, " ", "")
		_, err := w.Write(resp)
		if err != nil {
			http.Error(w, fmt.Errorf("internal server error: %v", err).Error(), http.StatusInternalServerError)
		}
	} else {
		data := NewResult(message, status)
		resp, _ := json.MarshalIndent(data, " ", "")
		_, err := w.Write(resp)
		if err != nil {
			http.Error(w, fmt.Errorf("internal server error: %v", err).Error(), http.StatusInternalServerError)
		}
	}

}
func (s *Server) DecodeJSON(r *http.Request) (*model.Event, error) {
	event := &model.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return nil, err
	}
	if event.UserID < 1 || event.EventID < 1 {
		return nil, errors.New("negative Id value")
	}

	return event, nil
}

func getParams(url *url.URL) (int, time.Time, error) {
	userId := url.Query().Get("user_id")
	date := url.Query().Get("date")

	id, err := strconv.Atoi(userId)
	if err != nil || id < 1 {
		if id < 1 {
			err = errors.New("negative userID")
		}
		return 0, time.Time{}, err
	}
	eventDate, err := parseDate(date)
	if err != nil {
		return 0, time.Time{}, err
	}
	return id, eventDate, nil
}
func parseDate(date string) (time.Time, error) {
	eventDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return time.Time{}, err
	}
	return eventDate, err
}
func methodError(w http.ResponseWriter, req, needle string) bool {
	if req != needle {
		http.Error(w, fmt.Sprintf("expected: %s, got:%s", needle, req), http.StatusInternalServerError)
	}
	return req != needle
}
