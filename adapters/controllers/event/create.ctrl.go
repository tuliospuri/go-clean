package event

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
    ctrl "tuliospuri/go-clean/adapters/controllers"
    bs "tuliospuri/go-clean/interactors"
    m "tuliospuri/go-clean/services/models"
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
    var params m.Generic
    body, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(body, &params)

    event := m.Event{
        Name: params["name"].(string),
    }

    // Execute interactor
    resp := c.eventCreateBs.Run(event)

    // Call presenter to format interactor's response
    jsonPresenter := pres.NewJsonPresenter()
    jsonPresenter.GetOutput(res, req, resp)
}
