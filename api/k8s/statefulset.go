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

func GetStatefulsets(ctx *gin.Context) {
	page, size := libs.GetPagination(ctx)
	name := ctx.Query("name")
	namespace := ctx.Query("namespace")
	client := utils.NewK8sConfig().InitClient()
	var continueToken string
	results := make([]*model.TStatefulSet, 0)
	for {
		namespaces, err := client.AppsV1().StatefulSets(namespace).List(context.Background(),
			metav1.ListOptions{Limit: 10, Continue: continueToken})
		if err != nil {
			libs.HttpServerError(ctx, "获取失败, err: %s", err.Error())
			return
		}
		for _, v := range namespaces.Items {
			results = append(results, &model.TStatefulSet{
				Id:                len(results) + 1,
				Name:              v.Name,
				Uid:               string(v.UID),
				CreateTime:        v.CreationTimestamp.Time,
				Namespace:         v.Namespace,
				Replicas:          v.Status.Replicas,
				UpdatedReplicas:   v.Status.UpdatedReplicas,
				CurrentReplicas:   v.Status.CurrentReplicas,
				AvailableReplicas: v.Status.AvailableReplicas,
			})
		}
		continueToken = namespaces.Continue
		if continueToken == "" {
			break
		}
	}
	// 根据名称过滤
	data := make([]*model.TStatefulSet, 0)
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
