package controllers

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/domain"
)

func SuccessResponse(code int, message string, data any) domain.Response {
	model := domain.Response{
		Status:     true,
		StatusCode: code,
		Message:    message,
		Error:      nil,
		Data:       data,
	}

	return model
}
func FailureResponse(code int, message string, err string) domain.Response {
	model := domain.Response{
		Status:     false,
		StatusCode: code,
		Message:    message,
		Error:      &err,
		Data:       nil,
	}

	return model
}
