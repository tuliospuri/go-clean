package main

import (
    "log"
    "net/http"

    rep "tuliospuri/go-clean/src/repositories"
    serv "tuliospuri/go-clean/src/services"
    eventBs "tuliospuri/go-clean/src/interactors/event"
    eventCtrl "tuliospuri/go-clean/src/adapters/controllers/event"
    val "tuliospuri/go-clean/src/adapters/validators"

    "github.com/gorilla/mux"
)

func main() {
    // Frameworks and drivers
    jsonSchema := val.NewJsonSchemaValidator()

    // Repositories
    eventRep := rep.NewEventRepository()
    errorRep := rep.NewErrorRepository()

    // Services
    eventServ := serv.NewEventService(eventRep)
    errorServ := serv.NewErrorService(errorRep)
    validatorServ := serv.NewValidatorService(jsonSchema)

    // Interactors
    eventCreateBs := eventBs.NewCreateBs(eventServ, errorServ, validatorServ)

    // Controllers
    eventCreateCtrl := eventCtrl.NewCreateController(eventCreateBs)

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/events", eventCreateCtrl.Execute).Methods("POST")
    log.Fatal(http.ListenAndServe(":8084", router))
}
