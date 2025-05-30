package usecaseimpls_test

import (
	"testing"

	"github.com/eduardo-gualberto/go.git/internal/application/gateways/usecaseimpls"
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/eventrepo"
)

func TestCreateEvent(t *testing.T) {
	testEventRepo := eventrepo.NewMockedEventRepo()
	uc := usecaseimpls.NewCreateEventImpl(testEventRepo)

	err := uc.Execute(0, 0, "RRULE:FREQ=DAILY")
	if err != nil {
		t.Errorf("expected error to be nil, found: %s\n", err.Error())
	}
}