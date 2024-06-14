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

func GetPods(ctx *gin.Context) {
	page, size := libs.GetPagination(ctx)
	ip := ctx.Query("ip")
	namespace := ctx.Query("namespace")
	client := utils.NewK8sConfig().InitClient()
	var continueToken string

	results := make([]*model.TPod, 0)
	for {
		pods, err := client.CoreV1().Pods(namespace).List(context.Background(),
			metav1.ListOptions{Limit: 10, Continue: continueToken})
		if err != nil {
			libs.HttpServerError(ctx, "获取失败, err: %w", err)
			return
		}
		for _, v := range pods.Items {
			results = append(results, &model.TPod{
				Id:            len(results) + 1,
				Name:          v.Name,
				Namespace:     v.Namespace,
				PodIp:         v.Status.PodIP,
				HostIp:        v.Status.HostIP,
				Status:        string(v.Status.Phase),
				CreateTime:    v.CreationTimestamp.Time,
				StartTime:     v.Status.StartTime.Time,
				Uid:           string(v.UID),
				RestartPolicy: string(v.Spec.RestartPolicy),
			})
		}
		if continueToken == "" {
			break
		}
	}

	// 根据地址过滤
	data := make([]*model.TPod, 0)
	if ip != "" {
		for _, v := range results {
			if strings.Contains(v.PodIp, ip) {
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
