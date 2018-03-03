package repositories

type eventRep struct {}

func NewEventRepository() EventRepository {
    return eventRep{}
}

func (r eventRep) Save(name string) int {
    return 2
}
