package presenters

import (
    "net/http"
    "encoding/json"
    m "tuliospuri/go-clean/src/services/models"
)

type jsonPresenter struct {}

func NewJsonPresenter() Presenter {
    return jsonPresenter{}
}

func (p jsonPresenter) GetOutput(res http.ResponseWriter, req *http.Request, params m.Generic) {
    res.
        Header().
        Set("Content-Type", "application/json; charset=UTF-8")

    if _, ok := params["code"]; ok {
        res.WriteHeader(http.StatusInternalServerError)
    } else {
        res.WriteHeader(http.StatusOK)
    }

    if err := json.NewEncoder(res).Encode(params); err != nil {
        panic(err)
    }
}
