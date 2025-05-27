package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/occurrencerepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
)

// ListOccurrencesImpl implements the ListOccurrences use case.
type ListOccurrencesImpl struct {
	repo occurrencerepo.OccurrenceRepo
}

// NewListOccurrencesImpl constructs a new ListOccurrencesImpl with the given OccurrenceRepo.
func NewListOccurrencesImpl(r occurrencerepo.OccurrenceRepo) usecases.ListOccurrences {
	return &ListOccurrencesImpl{repo: r}
}

// Execute retrieves all occurrences.
func (u *ListOccurrencesImpl) Execute() ([]*entities.OccurrenceEntity, error) {
	return u.repo.List()
}
