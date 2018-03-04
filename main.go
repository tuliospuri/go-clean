package main

import (
    "log"
    "net/http"

    rep "tuliospuri/go-clean/repositories"
    serv "tuliospuri/go-clean/services"
    eventBs "tuliospuri/go-clean/interactors/event"
    eventCtrl "tuliospuri/go-clean/adapters/controllers/event"

    "github.com/gorilla/mux"
)

func main() {
    // Repositories
    eventRep := rep.NewEventRepository()
    errorRep := rep.NewErrorRepository()

    // Services
    eventServ := serv.NewEventService(eventRep)
    errorServ := serv.NewErrorService(errorRep)

    // Interactors
    eventCreateBs := eventBs.NewCreateBs(eventServ, errorServ)

    // Controllers
    eventCreateCtrl := eventCtrl.NewCreateController(eventCreateBs)

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/events", eventCreateCtrl.Execute).Methods("POST")
    log.Fatal(http.ListenAndServe(":8084", router))
}
