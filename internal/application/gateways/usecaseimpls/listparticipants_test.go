package usecaseimpls_test

import (
	"testing"

	"github.com/eduardo-gualberto/go.git/internal/application/gateways/usecaseimpls"
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/participantrepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

func TestListParticipantsEmpty(t *testing.T) {
	testParticipantRepo := participantrepo.NewMockedParticipantRepo()
	uc := usecaseimpls.NewListParticipantsImpl(testParticipantRepo)

	list, err := uc.Execute()
	if err != nil {
		t.Errorf("expected error to be nil, found: %s\n", err.Error())
	}
	if len(list) != 0 {
		t.Errorf("expected list to be empty, got: %d\n", len(list))
	}
}

func TestListParticipantsNotEmpty(t *testing.T) {
	testParticipantRepo := participantrepo.NewMockedParticipantRepo()
	uc := usecaseimpls.NewListParticipantsImpl(testParticipantRepo)

	p0, err := testParticipantRepo.Create(entities.ParticipantEntity{ID: 0})
	if err != nil {
		panic(err)
	}

	list, err := uc.Execute()
	if err != nil {
		t.Errorf("expected error to be nil, found: %s\n", err.Error())
	}
	if len(list) != 1 {
		t.Errorf("expected list have a single element, got %d elements\n", len(list))
	}
	if list[0].ID != p0.ID {
		t.Errorf("expected only element of list to have ID = %d, got: %d\n", p0.ID, list[0].ID)
	}
}