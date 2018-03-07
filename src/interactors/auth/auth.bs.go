package auth

import (
    bs "tuliospuri/go-clean/src/interactors"
    serv "tuliospuri/go-clean/src/services"
    m "tuliospuri/go-clean/src/services/models"
)

type authBs struct {
    errorServ serv.ErrorService
}

func NewAuthBs(errorServ serv.ErrorService) bs.AuthBs {
    return authBs{errorServ}
}

func checkGroup(uri, group string) bool {
    if uri == "POST /events" && group == "admin" {
        return true
    }

    return false
}

func checkToken(token string) bool {
    if token == "123" {
        return true
    }

    return false
}

func (i authBs) Authorized(uri, group, token string) (bool, m.Generic) {

    okGroup := checkGroup(uri, group)
    okToken := checkToken(token)

    if okGroup || okToken {
        return true, m.Generic{}
    }

    return false, i.errorServ.Get("forbidden").Conv()
}

