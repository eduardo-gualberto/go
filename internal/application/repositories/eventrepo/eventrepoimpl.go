package eventrepo

import (
	"context"

	"github.com/eduardo-gualberto/go.git/internal/application/mappers"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"

	db "github.com/eduardo-gualberto/go.git/third_party/sqlc/gen"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// EventRepoImpl implements the EventRepo interface using a Postgres database.
type EventRepoImpl struct {
	q *db.Queries
}

// List retrieves all events.
func (r *EventRepoImpl) List() ([]*entities.EventEntity, error) {
	events, err := r.q.ListEvents(context.Background())
	if err != nil {
		return nil, err
	}
	return mappers.EventListDb2Entity(events), nil
}

// Create inserts a new event.
func (r *EventRepoImpl) Create(e entities.EventEntity) (*entities.EventEntity, error) {
	event, err := r.q.CreateEvent(context.Background(), db.CreateEventParams{
		ParticipantID: int64(e.ParticipantID),
		UserID:        int64(e.UserID),
		Rrule:         pgtype.Text{String: e.Rrule, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return mappers.EventDb2Entity(&event), nil
}

// NewEventRepoImpl creates a new EventRepo backed by Postgres.
func NewEventRepoImpl(conn *pgx.Conn) EventRepo {
	return &EventRepoImpl{q: db.New(conn)}
}