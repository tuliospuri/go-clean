package repositories

import (
    m "tuliospuri/go-clean/services/models"
)

type EventRepository interface {
    Save(name string) int
}

type ErrorRepository interface {
    GetByName(name string) m.Error
}
