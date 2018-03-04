package models

type Error struct {
    Code int
    Name string
    Message string
}

func (s Error) Conv() Generic {
    return Generic{
        "code": s.Code,
        "name": s.Name,
        "message": s.Message,
    }
}
