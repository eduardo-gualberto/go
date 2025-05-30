package usecaseimpls_test

import (
	"testing"

	"github.com/eduardo-gualberto/go.git/internal/application/gateways/usecaseimpls"
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/participantrepo"
)

func TestCreateParticipant(t *testing.T) {
	testParticipantRepo := participantrepo.NewMockedParticipantRepo()
	uc := usecaseimpls.NewCreateParticipantImpl(testParticipantRepo)

	err := uc.Execute("0011233334444", "teste participant", 0)
	if err != nil {
		t.Errorf("expected error to be nil, found: %s\n", err.Error())
	}
}