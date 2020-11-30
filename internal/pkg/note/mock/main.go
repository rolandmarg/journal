package mock

// TODO don't know if importing parent is good idea
import (
	"github.com/rolandmarg/jou/internal/pkg/note"
	"github.com/rolandmarg/jou/internal/pkg/random"
)

// Repository is mock implementation of note repository
type Repository struct {
	GetFn                 func(id int64) (*note.Note, error)
	GetInvoked            bool
	GetByJournalIDFn      func(id int64) ([]note.Note, error)
	GetByJournalIDInvoked bool
	CreateFn              func(journalID int64, title, body, mood string, tags []string) (int64, error)
	CreateInvoked         bool
	RemoveFn              func(id int64) error
	RemoveInvoked         bool
}

// Generate note struct with random data
func (r *Repository) Generate() *note.Note {
	n := &note.Note{}
	n.ID = random.Int64()
	n.JournalID = random.Int64()
	n.Title = random.String(128)
	n.Body = random.String(65535)
	n.Mood = random.String(64)
	n.Tags = random.Strings(64, 12)
	n.CreatedAt = random.Time()

	return n
}

// Get is mock implementation of note.Repository.Get
func (r *Repository) Get(id int64) (*note.Note, error) {
	r.GetInvoked = true
	return r.GetFn(id)
}

// GetByJournalID is mock implementation of journal.Repository.GetByJournalID
func (r *Repository) GetByJournalID(id int64) ([]note.Note, error) {
	r.GetByJournalIDInvoked = true
	return r.GetByJournalIDFn(id)
}

// Create is mock implementation of journal.Repository.Create
func (r *Repository) Create(journalID int64, title, body, mood string, tags []string) (int64, error) {
	r.CreateInvoked = true
	return r.CreateFn(journalID, title, body, mood, tags)
}

// Remove is mock implementation of note.Repository.Remove
func (r *Repository) Remove(id int64) error {
	r.RemoveInvoked = true
	return r.RemoveFn(id)
}
