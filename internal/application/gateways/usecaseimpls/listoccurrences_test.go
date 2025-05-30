package usecaseimpls_test

import (
	"testing"
	"time"

	"github.com/eduardo-gualberto/go.git/internal/application/gateways/usecaseimpls"
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/occurrencerepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

func TestListOccurrencesEmpty(t *testing.T) {
	testOccurrenceRepo := occurrencerepo.NewMockedOccurrenceRepo()
	uc := usecaseimpls.NewListOccurrencesImpl(testOccurrenceRepo)

	list, err := uc.Execute()
	if err != nil {
		t.Errorf("expected error to be nil, found: %s\n", err.Error())
	}
	if len(list) != 0 {
		t.Errorf("expected list to be empty, got: %d\n", len(list))
	}
}

func TestListOccurrencesNotEmpty(t *testing.T) {
	testOccurrenceRepo := occurrencerepo.NewMockedOccurrenceRepo()
	uc := usecaseimpls.NewListOccurrencesImpl(testOccurrenceRepo)

	o0, err := testOccurrenceRepo.Create(entities.OccurrenceEntity{ID: 0, OccursAt: time.Now()})
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
	if list[0].ID != o0.ID {
		t.Errorf("expected only element of list to have ID = %d, got: %d\n", o0.ID, list[0].ID)
	}
}