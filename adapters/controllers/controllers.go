package controllers

import (
    "net/http"
)

type Controller interface {
    Execute(res http.ResponseWriter, req *http.Request)
}
