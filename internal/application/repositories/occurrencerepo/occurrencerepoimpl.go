package occurrencerepo

import (
	"context"

	"github.com/eduardo-gualberto/go.git/internal/application/mappers"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"

	db "github.com/eduardo-gualberto/go.git/third_party/sqlc/gen"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// OccurrenceRepoImpl implements the OccurrenceRepo interface using a Postgres database.
type OccurrenceRepoImpl struct {
	q *db.Queries
}

// List retrieves all occurrences.
func (r *OccurrenceRepoImpl) List() ([]*entities.OccurrenceEntity, error) {
	occs, err := r.q.ListOccurrences(context.Background())
	if err != nil {
		return nil, err
	}
	return mappers.OccurrenceListDb2Entity(occs), nil
}

// Create inserts a new occurrence.
func (r *OccurrenceRepoImpl) Create(e entities.OccurrenceEntity) (*entities.OccurrenceEntity, error) {
	occ, err := r.q.CreateOccurrence(context.Background(), db.CreateOccurrenceParams{
		EventID:  int64(e.EventID),
		OccursAt: pgtype.Timestamptz{Time: e.OccursAt, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return mappers.OccurrenceDb2Entity(&occ), nil
}

// NewOccurrenceRepoImpl creates a new OccurrenceRepo backed by Postgres.
func NewOccurrenceRepoImpl(conn *pgx.Conn) OccurrenceRepo {
	return &OccurrenceRepoImpl{q: db.New(conn)}
}
