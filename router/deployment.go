package routers

import (
	"cicd/api/deployment"
	"github.com/gin-gonic/gin"
)

func DeploymentRouters(e *gin.RouterGroup) {
	resource := deployment.NewResourceTemplateHandler()
	app := deployment.NewAppHandler()
	healthcheck := deployment.NewHealthcheckTemplateHandler()
	e.GET("/deployment/resource_templates", resource.List)
	e.POST("/deployment/resource_template", resource.Create)
	e.PUT("/deployment/resource_template", resource.Update)
	e.DELETE("/deployment/resource_template", resource.Delete)

	e.GET("/deployment/healthcheck_templates", healthcheck.List)
	e.POST("/deployment/healthcheck_template", healthcheck.Create)
	e.PUT("/deployment/healthcheck_template", healthcheck.Update)
	e.DELETE("/deployment/healthcheck_template", healthcheck.Delete)

	e.GET("/deployment/apps", app.List)
	e.POST("/deployment/app", app.Create)
	e.PUT("/deployment/app", app.Update)
	e.DELETE("/deployment/app", app.Delete)

}
