package validator

type Validator interface {
	Process(interface{}) error
}
