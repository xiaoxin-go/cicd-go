package k8s

import (
	"cicd/libs"
	"cicd/model"
	"cicd/utils"
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func GetDeploys(ctx *gin.Context) {
	page, size := libs.GetPagination(ctx)
	name := ctx.Query("name")
	namespace := ctx.Query("namespace")
	client := utils.NewK8sConfig().InitClient()
	var continueToken string
	results := make([]*model.TDeploy, 0)
	for {
		namespaces, err := client.AppsV1().Deployments(namespace).List(context.Background(),
			metav1.ListOptions{Limit: 10, Continue: continueToken})
		if err != nil {
			libs.HttpServerError(ctx, "获取失败, err: %s", err.Error())
			return
		}
		for _, v := range namespaces.Items {
			results = append(results, &model.TDeploy{
				Id:                len(results) + 1,
				Name:              v.Name,
				Uid:               string(v.UID),
				CreateTime:        v.CreationTimestamp.Time,
				Namespace:         v.Namespace,
				Replicas:          v.Status.Replicas,
				UpdatedReplicas:   v.Status.UpdatedReplicas,
				ReadyReplicas:     v.Status.ReadyReplicas,
				AvailableReplicas: v.Status.AvailableReplicas,
			})
		}
		continueToken = namespaces.Continue
		if continueToken == "" {
			break
		}
	}
	// 根据名称过滤
	data := make([]*model.TDeploy, 0)
	if name != "" {
		for _, v := range results {
			if strings.Contains(v.Name, name) {
				data = append(data, v)
			}
		}
	} else {
		data = results
	}

	total := len(data)
	start := (page - 1) * size
	end := page * size
	if end > total {
		end = total
	}
	libs.HttpListSuccess(ctx, data[start:end], int64(total))
}
