package entities

type EventEntity struct {
	ID            int
	ParticipantID int
	UserID        int
	Rrule         string
}
