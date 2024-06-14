package routers

import (
	"cicd/api/k8s"
	"github.com/gin-gonic/gin"
)

func K8sRouters(e *gin.RouterGroup) {
	e.GET("/k8s/namespaces", k8s.GetNamespaces)
	e.POST("/k8s/namespace", k8s.CreateNamespace)
	e.PUT("/k8s/namespace", k8s.UpdateNamespace)
	e.DELETE("/k8s/namespace", k8s.DeleteNamespace)

	e.GET("/k8s/deploys", k8s.GetDeploys)
	e.GET("/k8s/statefulsets", k8s.GetStatefulsets)

	e.GET("/k8s/pods", k8s.GetPods)
	e.GET("/k8s/nodes", k8s.GetNodes)
	e.GET("/k8s/configmaps", k8s.GetConfigmaps)
	e.GET("/k8s/services", k8s.GetServices)
	e.GET("/k8s/secrets", k8s.GetSecrets)
}
