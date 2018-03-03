package presenters

import (
    "net/http"
    serv "tuliospuri/go-clean/services"
)

type Presenter interface {
    GetOutput(res http.ResponseWriter, req *http.Request, params serv.H)
}
