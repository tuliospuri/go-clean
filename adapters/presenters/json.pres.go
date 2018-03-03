package presenters

import (
    "net/http"
    "encoding/json"
    serv "tuliospuri/go-clean/services"
)

type jsonPresenter struct {}

func NewJsonPresenter() Presenter {
    return jsonPresenter{}
}

func (p jsonPresenter) GetOutput(res http.ResponseWriter, req *http.Request, params serv.H) {
    res.
        Header().
        Set("Content-Type", "application/json; charset=UTF-8")

    res.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(res).Encode(params); err != nil {
        panic(err)
    }
}
