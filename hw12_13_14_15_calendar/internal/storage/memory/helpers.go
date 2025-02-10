package memorystorage

import "github.com/google/uuid"

func (s *Storage) contains(id uuid.UUID) bool {
	_, ok := s.events[id]
	return ok
}
