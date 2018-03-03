package repositories

type EventRepository interface {
    Save(name string) int
}
