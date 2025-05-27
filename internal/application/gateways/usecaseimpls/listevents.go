package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/eventrepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
)

// ListEventsImpl implements the ListEvents use case.
type ListEventsImpl struct {
	repo eventrepo.EventRepo
}

// NewListEventsImpl constructs a new ListEventsImpl with the given EventRepo.
func NewListEventsImpl(r eventrepo.EventRepo) usecases.ListEvents {
	return &ListEventsImpl{repo: r}
}

// Execute retrieves all events.
func (u *ListEventsImpl) Execute() ([]*entities.EventEntity, error) {
	return u.repo.List()
}
