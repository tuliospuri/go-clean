package repositories

import (
    m "tuliospuri/go-clean/services/models"
)

type errorRep struct {
    errors []m.Error
}

func NewErrorRepository() ErrorRepository {
    errs := getErrors()
    return errorRep{errs}
}

func (r errorRep) GetByName(name string) m.Error {
    for _, v := range r.errors {
        if v.Name == name {
            return v
        }
    }
    return m.Error{}
}

func getErrors() []m.Error {
    errs := make([]m.Error, 0, 1)
    errs = append(errs, m.Error{1001, "invalid_name", "Invalid name"})
    return errs
}
