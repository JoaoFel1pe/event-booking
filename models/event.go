package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// add it to a database later
	events = append(events, e)

}
