package k8s

import (
	"cicd/libs"
	"cicd/model"
	"cicd/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func GetNodes(ctx *gin.Context) {
	page, size := libs.GetPagination(ctx)
	name := ctx.Query("name")
	client := utils.NewK8sConfig().InitClient()
	var continueToken string
	results := make([]*model.TNode, 0)
	for {
		namespaces, err := client.CoreV1().Nodes().List(context.Background(),
			metav1.ListOptions{Limit: 10, Continue: continueToken})
		if err != nil {
			libs.HttpServerError(ctx, "获取失败, err: %s", err.Error())
			return
		}
		for _, v := range namespaces.Items {
			results = append(results, &model.TNode{
				Id:         len(results) + 1,
				Name:       v.Name,
				Uid:        string(v.UID),
				Status:     string(v.Status.Phase),
				CreateTime: v.CreationTimestamp.Time,
				Address:    v.Status.Addresses[0].Address,
				Cpu:        fmt.Sprintf("%s", v.Status.Allocatable.Cpu()),
				Memory:     fmt.Sprintf("%s", v.Status.Allocatable.Memory()),
				PodCount:   fmt.Sprintf("%s", v.Status.Allocatable.Pods()),
			})
		}
		continueToken = namespaces.Continue
		if continueToken == "" {
			break
		}
	}
	// 根据名称过滤
	data := make([]*model.TNode, 0)
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
