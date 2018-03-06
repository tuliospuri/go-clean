package validators

import (
    "fmt"
    "github.com/xeipuuv/gojsonschema"
)

type jsonSchemaValidator struct {}

func NewJsonSchemaValidator() JsonSchemaValidator {
    return jsonSchemaValidator{}
}

func (v jsonSchemaValidator) Execute(json, against interface{}) bool {
    assets := "file:///go/src/tuliospuri/go-clean/assets/json/%s"
    file := fmt.Sprintf(assets, against.(string))
    schema := gojsonschema.NewReferenceLoader(file)
    document := gojsonschema.NewGoLoader(json)
    result, err := gojsonschema.Validate(schema, document)

    if err != nil {
        fmt.Println(err.Error())
        return false
    }

    if !result.Valid() {
        fmt.Println(result.Errors())
        return false
    }

    return true
}
