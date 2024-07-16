package deployment

import (
	"cicd/libs"
	"cicd/model"
	"fmt"
)

type HealthcheckTemplateHandler struct {
	libs.Controller
}

func NewHealthcheckTemplateHandler() *HealthcheckTemplateHandler {
	handler := &HealthcheckTemplateHandler{}
	handler.NewInstance = func() libs.Instance {
		return new(model.THealthcheckTemplate)
	}
	handler.NewResults = func() any {
		return &[]*model.THealthcheckTemplate{}
	}
	handler.Validate = func(data any) error {
		params := data.(*model.THealthcheckTemplate)
		fmt.Println(params)
		return nil
	}
	return handler
}
