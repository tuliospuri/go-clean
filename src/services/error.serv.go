package services

import (
    rep "tuliospuri/go-clean/src/repositories"
    m "tuliospuri/go-clean/src/services/models"
)

type errorService struct {
    errorRep rep.ErrorRepository
}

func NewErrorService(errorRep rep.ErrorRepository) ErrorService {
    return errorService{errorRep}
}

func (s errorService) Get(name string) m.Error {
    return s.errorRep.GetByName(name)
}
