package mappers

import (
	"github.com/eduardo-gualberto/go.git/core/entities"
	db "github.com/eduardo-gualberto/go.git/infra/db/gen"
)

func ParticipantDb2Entity(p *db.Participant) *entities.ParticipantEntity {
	return &entities.ParticipantEntity{
		ID:     int(p.ID),
		Number: p.ParticipantNumber,
		Name:   p.ParticipantName,
		UserID: int(p.UserID),
	}
}

func ParticipantListDb2Entity(ps []db.Participant) []*entities.ParticipantEntity {
	out := make([]*entities.ParticipantEntity, 0)
	for _, p := range ps {
		out = append(out, ParticipantDb2Entity(&p))
	}
	return out
}

func ParticipantEntity2Db(p *entities.ParticipantEntity) *db.Participant {
	return &db.Participant{
		ID:                int64(p.ID),
		ParticipantNumber: p.Number,
		ParticipantName:   p.Name,
		UserID:            int32(p.UserID),
	}
}

func ParticipantListEntity2Db(ps []entities.ParticipantEntity) []*db.Participant {
	out := make([]*db.Participant, len(ps))
	for _, p := range ps {
		out = append(out, ParticipantEntity2Db(&p))
	}
	return out
}
