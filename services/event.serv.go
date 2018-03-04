package services

import (
    rep "tuliospuri/go-clean/repositories"
    m "tuliospuri/go-clean/services/models"
)

type eventService struct {
    eventRep rep.EventRepository
}

func NewEventService(eventRep rep.EventRepository) EventService {
    return eventService{eventRep}
}

func (s eventService) Create(name string) m.Event {
    id := s.eventRep.Save(name)

    return m.Event{
        Id: id,
        Name: name,
    }
}
