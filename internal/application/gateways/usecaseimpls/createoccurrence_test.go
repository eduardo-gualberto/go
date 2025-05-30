package usecaseimpls_test

import (
	"testing"
	"time"

	"github.com/eduardo-gualberto/go.git/internal/application/gateways/usecaseimpls"
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/occurrencerepo"
)

func TestCreateOccurrence(t *testing.T) {
	testOccurrenceRepo := occurrencerepo.NewMockedOccurrenceRepo()
	uc := usecaseimpls.NewCreateOccurrenceImpl(testOccurrenceRepo)

	err := uc.Execute(0, time.Now())
	if err != nil {
		t.Errorf("expected error to be nil, found: %s\n", err.Error())
	}
}