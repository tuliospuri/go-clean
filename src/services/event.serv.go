package services

import (
    rep "tuliospuri/go-clean/src/repositories"
    m "tuliospuri/go-clean/src/services/models"
)

type eventService struct {
    eventRep rep.EventRepository
}

func NewEventService(eventRep rep.EventRepository) EventService {
    return eventService{eventRep}
}

func (s eventService) Create(event m.Event) m.Event {
    id := s.eventRep.Save(event.Name)

    return m.Event{
        Id: id,
        Name: event.Name,
    }
}
