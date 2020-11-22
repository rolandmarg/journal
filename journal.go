package main

// import (
// 	"fmt"
// )

// // Journal contains journal entries
// type Journal struct {
// 	name    string
// 	entries []*Entry
// }

// // NewJournal creates a new journal
// func NewJournal(name string) *Journal {
// 	return &Journal{name: name, entries: make([]*Entry, 0, 8)}
// }

// // AddEntry addes an entry to journal
// func (journal *Journal) AddEntry(input *EntryInput) *Entry {
// 	entry := NewEntry(input)

// 	journal.entries = append(journal.entries, entry)

// 	return entry
// }

// // RemoveEntry removes an entry from journal
// func (journal *Journal) RemoveEntry(id int) *Entry {
// 	for _, entry := range journal.entries {
// 		if entry.id == id {
// 			return entry.SoftDeleteEntry()
// 		}
// 	}

// 	return nil
// }

// func (journal *Journal) String() string {
// 	str := fmt.Sprintln("journal:")

// 	for _, e := range journal.entries {
// 		if e.IsEntryDeleted() {
// 			continue
// 		}
// 		str = fmt.Sprintln(str, " entry:")
// 		str = fmt.Sprintln(str, "   id:", e.id)
// 		str = fmt.Sprintln(str, "   title:", e.title)
// 		str = fmt.Sprintln(str, "   body:", e.body)
// 		if e.mood != "" {
// 			str = fmt.Sprintln(str, "   mood:", e.mood)
// 		}
// 		if e.tags != nil {
// 			str = fmt.Sprintln(str, "   tags:", e.tags)
// 		}
// 		str = fmt.Sprintln(str, "   createdAt:", e.createdAt.Format("2006-01-02 15:04:05"))
// 	}

// 	return str
// }
