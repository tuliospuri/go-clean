package interactors

import (
    m "tuliospuri/go-clean/src/services/models"
)

type EventCreateBs interface {
    Run(event m.Generic) m.Generic
}

type AuthBs interface {
    Authorized(uri, group, token string) (bool, m.Generic)
}
