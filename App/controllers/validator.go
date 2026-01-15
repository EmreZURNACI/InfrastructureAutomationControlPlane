package controllers

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

func Validate(model interface{}) error {
	validate := validator.New()
	if err := validate.Struct(model); err != nil {
		return err
	}
	return nil
}
func ParseToModel(params []byte, model interface{}) (interface{}, error) {
	if err := json.Unmarshal(params, model); err != nil {
		return nil, err
	}
	return model, nil
}
