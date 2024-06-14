package deployment

import (
	"cicd/libs"
	"cicd/model"
	"fmt"
)

type ResourceTemplateHandler struct {
	libs.Controller
}

func NewResourceTemplateHandler() *ResourceTemplateHandler {
	handler := &ResourceTemplateHandler{}
	handler.NewInstance = func() libs.Instance {
		return new(model.TResourceTemplate)
	}
	handler.NewResults = func() any {
		return &[]*model.TResourceTemplate{}
	}
	handler.Validate = func(data any) error {
		resourceTemplate := data.(model.TResourceTemplate)
		if resourceTemplate.MemoryUnit != model.MemoryUnitMi && resourceTemplate.MemoryUnit != model.MemoryUnitGi {
			return fmt.Errorf("memory_unit must be %s or %s, not %s", model.MemoryUnitMi, model.MemoryUnitGi, resourceTemplate.MemoryUnit)
		}
		return nil
	}
	return handler
}
