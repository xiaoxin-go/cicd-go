package model

type TConfigTemplate struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name"`
	Containers  string `gorm:""`
	Annotations string `gorm:"column:annotations" json:"annotations"`
	Labels      string `gorm:"column:labels" json:"labels"`
	Desc        string `gorm:"desc" json:"desc"`
}

type MemoryUnitType string

const (
	MemoryUnitMi MemoryUnitType = "Mi"
	MemoryUnitGi MemoryUnitType = "Gi"
)

type TResourceTemplate struct {
	BaseModel
	Name        string         `gorm:"column:name" json:"name"`
	Cpu         float32        `gorm:"column:cpu" json:"cpu"`
	Memory      int            `gorm:"memory" json:"memory"`
	MemoryUnit  MemoryUnitType `gorm:"memory_unit" json:"memory_unit"`
	Description string         `gorm:"description" json:"description"`
}

type THealthcheckTemplate struct {
	BaseModel
	Name           string `gorm:"column:name" json:"name"`
	ReadinessProbe string `json:"readiness_probe"`
	LivenessProbe  string `json:"liveness_probe"`
	StartupProbe   string `json:"startup_probe"`
	Desc           string `gorm:"desc" json:"desc"`
}
