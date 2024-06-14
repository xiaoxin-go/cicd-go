package pkg

import "cicd/model"

type AppServiceInterface interface {
	Deploy()
	Scale()
	Restart()
}

func NewAppService(serviceId int) AppServiceInterface {
	return &appServiceHandler{}
}

type appServiceHandler struct {
	data                *model.TService
	app                 *model.TApp
	healthcheckTemplate *model.THealthcheckTemplate
	configTemplate      *model.TConfigTemplate
	resourceTemplate    *model.TResourceTemplate
}

func (h *appServiceHandler) Deploy() {

}
func (h *appServiceHandler) Scale() {

}
func (h *appServiceHandler) Restart() {

}
