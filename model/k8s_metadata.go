package model

import (
	v1 "k8s.io/api/core/v1"
	"time"
)

type TPod struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	PodIp         string    `json:"pod_ip"`
	HostIp        string    `json:"host_ip"`
	Status        string    `json:"status"`
	Namespace     string    `json:"namespace"`
	RestartPolicy string    `json:"restart_policy"`
	CreateTime    time.Time `json:"create_time"`
	StartTime     time.Time `json:"start_time"`
	NodeName      string    `json:"node_name"`
	Uid           string    `json:"uid"`
}

type TNode struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Uid        string    `json:"uid"`
	CreateTime time.Time `json:"create_time"`
	Status     string    `json:""`
	Cpu        string    `json:"cpu"`
	Memory     string    `json:"memory"`
	PodCount   string    `json:"pod_count"`
	Address    string    `json:"address"`
}

type TDeploy struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Uid               string    `json:"uid"`
	Namespace         string    `json:"namespace"`
	Replicas          int32     `json:"replicas"`
	UpdatedReplicas   int32     `json:"updated_replicas"`
	ReadyReplicas     int32     `json:"ready_replicas"`
	AvailableReplicas int32     `json:"availableReplicas"`
	CreateTime        time.Time `json:"create_time"`
}

type TStatefulSet struct {
	Id                int       `json:"id"`
	Uid               string    `json:"uid"`
	Name              string    `json:"name"`
	Namespace         string    `json:"namespace"`
	CreateTime        time.Time `json:"create_time"`
	Replicas          int32     `json:"replicas"`
	UpdatedReplicas   int32     `json:"updated_replicas"`
	CurrentReplicas   int32     `json:"current_replicas"`
	AvailableReplicas int32     `json:"available_replicas"`
}

type TService struct {
	Id         int              `json:"id"`
	Uid        string           `json:"uid"`
	Name       string           `json:"name"`
	Namespace  string           `json:"namespace"`
	CreateTime time.Time        `json:"create_time"`
	ClusterIp  string           `json:"cluster_ip"`
	Type       string           `json:"type"`
	Ports      []v1.ServicePort `json:"ports"`
}

type TConfigmap struct {
	Id         int               `json:"id"`
	Uid        string            `json:"uid"`
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	CreateTime time.Time         `json:"create_time"`
	Data       map[string]string `json:"data"`
}

type TSecret struct {
	Id         int               `json:"id"`
	Uid        string            `json:"uid"`
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	CreateTime time.Time         `json:"create_time"`
	Type       string            `json:"type"`
	Data       map[string][]byte `json:"data"`
}
