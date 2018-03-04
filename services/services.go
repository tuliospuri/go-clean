package services

import (
    m "tuliospuri/go-clean/services/models"
)

type EventService interface {
    Create(name string) m.Event
}

type ErrorService interface {
    Get(name string) m.Error
}
