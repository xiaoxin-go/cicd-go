package deployment

import (
	"cicd/libs"
	"cicd/model"
	"fmt"
)

type AppHandler struct {
	libs.Controller
}

func NewAppHandler() *AppHandler {
	handler := &AppHandler{}
	handler.NewInstance = func() libs.Instance {
		return new(model.TApp)
	}
	handler.NewResults = func() any {
		return &[]*model.TApp{}
	}
	handler.Validate = func(data any) error {
		app := data.(model.TApp)
		if app.ServiceType != model.ServiceTypeStateless && app.ServiceType != model.ServiceTypeStated {
			return fmt.Errorf("service_type must be %s or %s, not %s", model.ServiceTypeStated, model.ServiceTypeStateless, app.ServiceType)
		}
		return nil
	}
	return handler
}
