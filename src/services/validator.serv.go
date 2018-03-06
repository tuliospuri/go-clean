package services

import (
    m "tuliospuri/go-clean/src/services/models"
    val "tuliospuri/go-clean/src/adapters/validators"
)

type validatorService struct {
    jsonSchema val.JsonSchemaValidator
}

func NewValidatorService(jsonSchema val.JsonSchemaValidator) ValidatorService {
    return validatorService{jsonSchema}
}

func (s validatorService) IsValid(params m.Generic, schema string) bool {
    return s.jsonSchema.Execute(params, schema)
}
