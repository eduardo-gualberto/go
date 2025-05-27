package mappers

import (
	"github.com/eduardo-gualberto/go.git/internal/core/entities"

	db "github.com/eduardo-gualberto/go.git/third_party/sqlc/gen"
	"github.com/jackc/pgx/v5/pgtype"
)

// OccurrenceDb2Entity maps a db.Occurrence to a domain OccurrenceEntity.
func OccurrenceDb2Entity(o *db.Occurrence) *entities.OccurrenceEntity {
	return &entities.OccurrenceEntity{
		ID:       int(o.ID),
		EventID:  int(o.EventID),
		OccursAt: o.OccursAt.Time,
	}
}

// OccurrenceListDb2Entity maps a slice of db.Occurrence to a slice of domain OccurrenceEntity pointers.
func OccurrenceListDb2Entity(os []db.Occurrence) []*entities.OccurrenceEntity {
	out := make([]*entities.OccurrenceEntity, 0)
	for _, o := range os {
		out = append(out, OccurrenceDb2Entity(&o))
	}
	return out
}

// OccurrenceEntity2Db maps a domain OccurrenceEntity to a db.Occurrence.
func OccurrenceEntity2Db(o *entities.OccurrenceEntity) *db.Occurrence {
	return &db.Occurrence{
		ID:       int64(o.ID),
		EventID:  int64(o.EventID),
		OccursAt: pgtype.Timestamptz{Time: o.OccursAt, Valid: true},
	}
}

// OccurrenceListEntity2Db maps a slice of domain OccurrenceEntity to a slice of db.Occurrence pointers.
func OccurrenceListEntity2Db(os []entities.OccurrenceEntity) []*db.Occurrence {
	out := make([]*db.Occurrence, len(os))
	for _, o := range os {
		out = append(out, OccurrenceEntity2Db(&o))
	}
	return out
}
