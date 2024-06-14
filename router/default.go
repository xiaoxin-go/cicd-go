package routers

import (
	"github.com/gin-gonic/gin"
)

type Option func(engine *gin.RouterGroup)

var options = make([]Option, 0)

func Include(opts ...Option) {
	options = append(options, opts...)
}

func IncludeRouter() {
	Include(K8sRouters)
	Include(DeploymentRouters)
}

// Init 初始化
func Init(r *gin.RouterGroup) *gin.RouterGroup {
	IncludeRouter()
	for _, opt := range options {
		opt(r)
	}
	return r
}
