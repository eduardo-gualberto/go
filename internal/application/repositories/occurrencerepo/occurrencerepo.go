package occurrencerepo

import (
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

// OccurrenceRepo defines the interface for occurrence persistence operations.
type OccurrenceRepo interface {
	List() ([]*entities.OccurrenceEntity, error)
	Create(e entities.OccurrenceEntity) (*entities.OccurrenceEntity, error)
}