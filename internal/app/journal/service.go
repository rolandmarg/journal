package journal

import (
	"fmt"

	"github.com/rolandmarg/jou/internal/pkg/journal"
	"github.com/rolandmarg/jou/internal/pkg/note"
)

// Service provides journal functions
type Service struct {
	j journal.Repository
	n note.Repository
	// TODO add custom logging for errors info and so on
}

// MakeService creates journal service
func MakeService(j journal.Repository, n note.Repository) *Service {
	s := &Service{j, n}

	return s
}

// Get journal by name
func (s *Service) Get(name string) (*journal.Journal, error) {
	j, err := s.j.Get(name)
	if err != nil {
		return nil, fmt.Errorf(`Could not get journal: %w`, err)
	}

	j.Notes, err = s.n.GetByJournalID(j.ID)
	if err != nil {
		return j, fmt.Errorf(`Could not get journal notes: %w`, err)
	}

	return j, nil
}

// GetAll journals
func (s *Service) GetAll() ([]journal.Journal, error) {
	journals, err := s.j.GetAll()
	if err != nil {
		return nil, fmt.Errorf(`Could not get journals: %w`, err)
	}

	for _, j := range journals {
		j.Notes, err = s.n.GetByJournalID(j.ID)
		if err != nil {
			// TODO maybe try getting other journal notes instead of return
			return journals, fmt.Errorf(`Could not get journal "%v" notes: %w`, j.Name, err)
		}
	}

	return journals, nil
}

// Create journal
func (s *Service) Create(name string, isDefault bool) error {
	_, err := s.j.Create(name)
	if err != nil {
		return fmt.Errorf(`Could not create journal: %w`, err)
	}

	if isDefault {
		err = s.j.SetDefault(name)
		if err != nil {
			return fmt.Errorf(`Journal created, but could not set default: %w`, err)
		}
	}

	return nil
}
