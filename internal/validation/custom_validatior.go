package validation

type CustomValidator struct{}

func (v *CustomValidator) ValidateStruct(obj interface{}) error {
	return Provider.Validator().Struct(obj)
}

func (v *CustomValidator) Engine() interface{} {
	return Provider.Validator()
}
