package event

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
    ctrl "tuliospuri/go-clean/adapters/controllers"
    bs "tuliospuri/go-clean/interactors"
    serv "tuliospuri/go-clean/services"
    pres "tuliospuri/go-clean/adapters/presenters"
)

type eventCreateCtrl struct {
    eventCreateBs bs.EventCreateBs
}

func NewCreateController(eventCreateBs bs.EventCreateBs) ctrl.Controller {
    return eventCreateCtrl{eventCreateBs}
}

func (c eventCreateCtrl) Execute(res http.ResponseWriter, req *http.Request) {

    // Extracts data from http request and converts to the interactor
    var params serv.H
    body, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(body, &params)

    event := serv.Event{
        Name: params["name"].(string),
    }

    // Execute interactor
    resp := c.eventCreateBs.Run(event)

    // Call presenter to format interactor's response
    jsonPresenter := pres.NewJsonPresenter()
    jsonPresenter.GetOutput(res, req, resp)
}