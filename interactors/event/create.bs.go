package event

import (
    bs "tuliospuri/go-clean/interactors"
    serv "tuliospuri/go-clean/services"
    m "tuliospuri/go-clean/services/models"
)

type eventCreateBs struct {
    eventServ serv.EventService
    errorServ serv.ErrorService
}

func NewCreateBs(
    eventServ serv.EventService,
    errorServ serv.ErrorService) bs.EventCreateBs {

    return eventCreateBs{eventServ, errorServ}
}

func (i eventCreateBs) Run(event m.Event) m.Generic {

    if event.Name == "test" {
        return i.errorServ.Get("invalid_name").Conv()
    }

    created := i.eventServ.Create(event.Name)

    return m.Generic{
        "Id": created.Id,
        "Name": event.Name,
    }
}
