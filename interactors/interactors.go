package interactors

import (
    serv "tuliospuri/go-clean/services"
)

type EventCreateBs interface {
    Run(event serv.Event) serv.H
}
