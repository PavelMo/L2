package storage

import (
	"dev11/internal/model"
	"errors"
	"fmt"
	"sync"
	"time"
)

type DB struct {
	sync.RWMutex
	data map[string]model.Event
}

func NewDB() *DB {
	return &DB{
		data: make(map[string]model.Event),
	}
}

func (d *DB) CreateEvent(event *model.Event) error {
	id := fmt.Sprintf("%d%d", event.UserID, event.EventID)
	d.Lock()
	if _, ok := d.data[id]; ok {
		d.Unlock()
		return errors.New("event with this id is currently exists")
	}
	d.data[id] = *event
	d.Unlock()

	return nil
}
func (d *DB) UpdateEvent(event *model.Event) error {
	id := fmt.Sprintf("%d%d", event.UserID, event.EventID)
	d.Lock()
	if _, ok := d.data[id]; !ok {
		d.Unlock()
		return errors.New("no events found with this id")
	}
	d.Unlock()
	d.data[id] = *event
	return nil
}
func (d *DB) DeleteEvent(event *model.Event) {
	id := fmt.Sprintf("%d%d", event.UserID, event.EventID)
	d.Lock()
	delete(d.data, id)
	d.Unlock()
}
func (d *DB) EventsForDay(userId int, date time.Time) ([]model.Event, error) {
	var events []model.Event
	year, month, day := date.Date()
	d.RLock()
	for _, event := range d.data {
		eventYear, eventMonth, eventDay := event.Date.Date()
		if eventYear == year && eventMonth == month && eventDay == day && event.UserID == userId {
			events = append(events, event)
		}
	}
	d.RUnlock()
	if len(events) < 1 {
		return nil, errors.New("no data found for this day")
	}
	return events, nil
}
func (d *DB) EventsForWeek(userId int, date time.Time) ([]model.Event, error) {
	var events []model.Event
	year, week := date.ISOWeek()
	d.RLock()
	for _, event := range d.data {
		eventYear, eventWeek := event.Date.ISOWeek()
		if eventYear == year && eventWeek == week && event.UserID == userId {
			events = append(events, event)
		}
	}
	d.RUnlock()
	if len(events) < 1 {
		return nil, errors.New("no data found for this week")
	}
	return events, nil
}
func (d *DB) EventsForMonth(userId int, date time.Time) ([]model.Event, error) {
	var events []model.Event
	year, month, _ := date.Date()
	d.RLock()
	for _, event := range d.data {
		eventYear, eventMonth, _ := event.Date.Date()
		if eventYear == year && eventMonth == month && event.UserID == userId {
			events = append(events, event)
		}
	}
	d.RUnlock()
	if len(events) < 1 {
		return nil, errors.New("no data found for this month")
	}
	return events, nil
}
