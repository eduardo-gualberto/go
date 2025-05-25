package participantrepo

import (
	"context"

	"github.com/eduardo-gualberto/go.git/internal/application/mappers"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	db "github.com/eduardo-gualberto/go.git/third_party/sqlc/gen"
	"github.com/jackc/pgx/v5"
)

type ParticipantRepoImpl struct {
	q *db.Queries
}

func (r *ParticipantRepoImpl) List() ([]*entities.ParticipantEntity, error) {
	participants, err := r.q.ListParticipants(context.Background())
	if err != nil {
		return nil, err
	}
	return mappers.ParticipantListDb2Entity(participants), nil
}

func (r *ParticipantRepoImpl) GetByID(id int) (*entities.ParticipantEntity, error) {
	participant, err := r.q.GetParticipantByID(context.Background(), int64(id))
	if err != nil {
		return nil, err
	}
	return mappers.ParticipantDb2Entity(&participant), nil
}

func (r *ParticipantRepoImpl) Create(e entities.ParticipantEntity) (*entities.ParticipantEntity, error) {
	participant, err := r.q.CreateParticipant(context.Background(), db.CreateParticipantParams{
		ParticipantNumber: e.Number,
		ParticipantName:   e.Name,
		UserID:            int32(e.UserID),
	})
	if err != nil {
		return nil, err
	}
	return mappers.ParticipantDb2Entity(&participant), nil
}

func NewParticipantRepoImpl(conn *pgx.Conn) ParticipantRepo {
	return &ParticipantRepoImpl{
		q: db.New(conn),
	}
}
