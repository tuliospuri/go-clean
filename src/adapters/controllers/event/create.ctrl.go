package event

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
    ctrl "tuliospuri/go-clean/src/adapters/controllers"
    bs "tuliospuri/go-clean/src/interactors"
    m "tuliospuri/go-clean/src/services/models"
    pres "tuliospuri/go-clean/src/adapters/presenters"
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

    // Execute interactor
    resp := c.eventCreateBs.Run(params)

    // Call presenter to format interactor's response
    jsonPres := pres.NewJsonPresenter()
    jsonPres.GetOutput(res, req, resp)
}
