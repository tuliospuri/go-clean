package event

import (
    "tuliospuri/go-clean/interactors"
    serv "tuliospuri/go-clean/services"
)

type eventCreateBs struct {
    eventServ serv.EventService
}

func NewCreateBs(eventServ serv.EventService) interactors.EventCreateBs {
    return eventCreateBs{eventServ}
}

func (i eventCreateBs) Run(event serv.Event) serv.H {
    created := i.eventServ.Create(event.Name)

    return serv.H{
        "Id": created.Id,
        "Name": event.Name,
    }
}
