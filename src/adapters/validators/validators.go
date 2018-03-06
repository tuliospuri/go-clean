package validators

type JsonSchemaValidator interface {
    Execute(json, against interface{}) bool
}
