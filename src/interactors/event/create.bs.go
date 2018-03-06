package event

import (
    bs "tuliospuri/go-clean/src/interactors"
    serv "tuliospuri/go-clean/src/services"
    m "tuliospuri/go-clean/src/services/models"
)

type eventCreateBs struct {
    eventServ serv.EventService
    errorServ serv.ErrorService
    validatorServ serv.ValidatorService
}

func NewCreateBs(
    eventServ serv.EventService,
    errorServ serv.ErrorService,
    validatorServ serv.ValidatorService) bs.EventCreateBs {

    return eventCreateBs{eventServ, errorServ, validatorServ}
}

func (i eventCreateBs) Run(event m.Generic) m.Generic {

    if !i.validatorServ.IsValid(event, "event.json") {
        return i.errorServ.Get("invalid_schema").Conv()
    }

    if event["name"] == "test" {
        return i.errorServ.Get("invalid_name").Conv()
    }

    newEvent := m.Event{
        Name: event["name"].(string),
    }

    created := i.eventServ.Create(newEvent)

    return created.Conv()
}
