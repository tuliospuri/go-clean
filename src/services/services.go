package services

import (
    m "tuliospuri/go-clean/src/services/models"
)

type EventService interface {
    Create(event m.Event) m.Event
}

type ErrorService interface {
    Get(name string) m.Error
}

type ValidatorService interface {
    IsValid(params m.Generic, schema string) bool
}
