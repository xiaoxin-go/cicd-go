package k8s

import (
	"cicd/libs"
	"cicd/model"
	"cicd/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type NamespaceController struct {
}

type createNamespaceParams struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

func CreateNamespace(ctx *gin.Context) {
	params := &createNamespaceParams{}
	if e := ctx.ShouldBindJSON(params); e != nil {
		libs.HttpParamsError(ctx, "parse params failed, %s", e.Error())
		return
	}
	client := utils.NewK8sConfig().InitClient()
	namespace := &v1.Namespace{}
	namespace.Name = params.Name
	namespace.Labels = params.Labels
	opts := metav1.CreateOptions{}
	result, e := client.CoreV1().Namespaces().Create(context.Background(), namespace, opts)
	if e != nil {
		libs.HttpServerError(ctx, "create namespace failed, %s", e.Error())
		return
	}
	item := &model.TNamespace{
		Name:   result.Name,
		Uid:    string(result.UID),
		Status: string(result.Status.Phase),
	}
	item.CreatedAt = result.CreationTimestamp.Time
	libs.HttpSuccess(ctx, item, "创建成功")
}

func DeleteNamespace(ctx *gin.Context) {
	name := ctx.Query("name")
	client := utils.NewK8sConfig().InitClient()
	if e := client.CoreV1().Namespaces().Delete(context.Background(), name, metav1.DeleteOptions{}); e != nil {
		libs.HttpServerError(ctx, "delete namespace failed, %s", e.Error())
		return
	}
	libs.HttpSuccess(ctx, name, "%s deleted", name)
}

func UpdateNamespace(ctx *gin.Context) {
	params := &createNamespaceParams{}
	if e := ctx.ShouldBindJSON(params); e != nil {
		libs.HttpParamsError(ctx, "parse params failed, %s", e.Error())
		return
	}
	client := utils.NewK8sConfig().InitClient()
	ns := &v1.Namespace{}
	ns.Name = params.Name
	ns.Labels = params.Labels
	opts := metav1.UpdateOptions{}
	result, e := client.CoreV1().Namespaces().Update(context.Background(), ns, opts)
	if e != nil {
		libs.HttpServerError(ctx, "update namespace failed, %s", e.Error())
		return
	}
	libs.HttpSuccess(ctx, result, "update success")
}

func GetNamespaces(ctx *gin.Context) {
	page, size := libs.GetPagination(ctx)
	name := ctx.Query("name")
	client := utils.NewK8sConfig().InitClient()
	var continueToken string
	results := make([]*model.TNamespace, 0)
	for {
		namespaces, err := client.CoreV1().Namespaces().List(context.Background(),
			metav1.ListOptions{Limit: 10, Continue: continueToken})
		if err != nil {
			libs.HttpServerError(ctx, "获取失败, err: %s", err.Error())
			return
		}
		for _, v := range namespaces.Items {
			item := &model.TNamespace{
				Name:   v.Name,
				Uid:    string(v.UID),
				Status: string(v.Status.Phase),
			}
			item.Id = len(results) + 1
			item.CreatedAt = v.CreationTimestamp.Time
			bs, _ := json.Marshal(v.Labels)
			item.Labels = string(bs)
			results = append(results, item)
		}
		continueToken = namespaces.Continue
		if continueToken == "" {
			break
		}
	}
	// 根据名称过滤
	data := make([]*model.TNamespace, 0)
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
