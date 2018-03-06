package presenters

import (
    "net/http"
    m "tuliospuri/go-clean/src/services/models"
)

type Presenter interface {
    GetOutput(res http.ResponseWriter, req *http.Request, params m.Generic)
}
