package mappers

import (
   "github.com/eduardo-gualberto/go.git/internal/core/entities"

   db "github.com/eduardo-gualberto/go.git/third_party/sqlc/gen"
   "github.com/jackc/pgx/v5/pgtype"
)

// EventDb2Entity maps a db.Event to a domain EventEntity.
func EventDb2Entity(e *db.Event) *entities.EventEntity {
   return &entities.EventEntity{
       ID:            int(e.ID),
       ParticipantID: int(e.ParticipantID),
       UserID:        int(e.UserID),
       Rrule:         e.Rrule.String,
   }
}

// EventListDb2Entity maps a slice of db.Event to a slice of domain EventEntity pointers.
func EventListDb2Entity(es []db.Event) []*entities.EventEntity {
   out := make([]*entities.EventEntity, 0)
   for _, e := range es {
       out = append(out, EventDb2Entity(&e))
   }
   return out
}

// EventEntity2Db maps a domain EventEntity to a db.Event.
func EventEntity2Db(e *entities.EventEntity) *db.Event {
   return &db.Event{
       ID:            int64(e.ID),
       ParticipantID: int64(e.ParticipantID),
       UserID:        int64(e.UserID),
       Rrule:         pgtype.Text{String: e.Rrule, Valid: true},
   }
}

// EventListEntity2Db maps a slice of domain EventEntity to a slice of db.Event pointers.
func EventListEntity2Db(es []entities.EventEntity) []*db.Event {
   out := make([]*db.Event, len(es))
   for _, e := range es {
       out = append(out, EventEntity2Db(&e))
   }
   return out
}