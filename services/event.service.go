package services

import (
    rep "tuliospuri/go-clean/repositories"
)

type eventService struct {
    eventRep rep.EventRepository
}

func NewEventService(eventRep rep.EventRepository) EventService {
    return eventService{eventRep}
}

func (s eventService) Create(name string) Event {
    id := s.eventRep.Save(name)

    return Event{
        Id: id,
        Name: name,
    }
}
