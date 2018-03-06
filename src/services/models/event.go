package models

type Event struct {
    Id int
    Name string
}

func (s Event) Conv() Generic {
    return Generic{
        "id": s.Id,
        "name": s.Name,
    }
}
