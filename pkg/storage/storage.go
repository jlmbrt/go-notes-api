package storage

import "fmt"

type Note struct {
	Message string `json:"message" binding:"required,min=1"`
	Tag     string `json:"tag"`
}

type Storage struct {
	db         []Note
	tagIndexed map[string][]*Note
}

// Create a new storage object with his own DB
func New() *Storage {
	return &Storage{
		db:         make([]Note, 0),
		tagIndexed: make(map[string][]*Note),
	}
}

// Add a new note in DB
func (s *Storage) Create(note Note) {
	// store new item
	s.db = append(s.db, note)

	// if no tag, work is done
	if note.Tag == "" {
		return
	}

	// index item with his tag
	if _, ok := s.tagIndexed[note.Tag]; !ok {
		fmt.Println("Create index")
		s.tagIndexed[note.Tag] = make([]*Note, 0)
	}

	index := s.tagIndexed[note.Tag]

	// store *ptr only, not the full object
	index = append(index, &s.db[len(s.db)-1])

	s.tagIndexed[note.Tag] = index
}

// Return All item in DB
func (s *Storage) FinAll() []Note {
	return s.db
}

// Return all item listed in tagIndexed[<tag>]
func (s *Storage) FindWithTag(tag string) []Note {
	result := make([]Note, 0)

	index, ok := s.tagIndexed[tag]

	if !ok {
		return result
	}

	for _, note := range index {
		result = append(result, *note)
	}

	return result
}
