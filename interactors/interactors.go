package interactors

import (
    m "tuliospuri/go-clean/services/models"
)

type EventCreateBs interface {
    Run(event m.Event) m.Generic
}
