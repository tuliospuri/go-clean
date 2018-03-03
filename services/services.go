package services

type EventService interface {
    Create(name string) Event
}
