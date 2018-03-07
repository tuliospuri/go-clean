package main

import (
    "log"
    "net/http"

    rep "tuliospuri/go-clean/src/repositories"
    serv "tuliospuri/go-clean/src/services"
    eventBs "tuliospuri/go-clean/src/interactors/event"
    authBs "tuliospuri/go-clean/src/interactors/auth"
    eventCtrl "tuliospuri/go-clean/src/adapters/controllers/event"
    val "tuliospuri/go-clean/src/adapters/validators"
    mw "tuliospuri/go-clean/src/adapters/middlewares"

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
    authBs := authBs.NewAuthBs(errorServ)
    eventCreateBs := eventBs.NewCreateBs(eventServ, errorServ, validatorServ)

    // Controllers and middlewares
    authMw := mw.NewAuthMiddleware(authBs)
    eventCreateCtrl := eventCtrl.NewCreateController(eventCreateBs)

    router := mux.NewRouter().StrictSlash(true)

    router.
        Use(authMw.Execute)
    router.
        HandleFunc("/events", eventCreateCtrl.Execute).Methods("POST")

    log.Fatal(http.ListenAndServe(":8084", router))
}
