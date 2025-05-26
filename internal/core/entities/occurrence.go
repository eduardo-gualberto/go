package entities

import "time"

type OccurrenceEntity struct {
	ID       int
	EventID  int
	OccursAt time.Time
}
