package services

import "github.com/n1jke/oop-bsuir-2025/lr-1/internal/domain"

type EventService struct {
	store EventStorage
}

func NewEventService(store EventStorage) *EventService {
	return &EventService{store: store}
}

func (s *EventService) Publish(e domain.Event) error {
	return s.store.Save(e)
}

func (s *EventService) QueryAll() []domain.Event {
	return s.store.QueryAll()
}
