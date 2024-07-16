package model

type TConfigTemplate struct {
	BaseModel
	Name        string                      `gorm:"column:name" json:"name"`
	Containers  []*TConfigTemplateContainer `json:"containers"`
	Annotations string                      `gorm:"column:annotations" json:"annotations"`
	Labels      string                      `gorm:"column:labels" json:"labels"`
	Description string                      `gorm:"description" json:"description"`
}

type TConfigTemplateContainer struct {
	BaseModel
	ConfigTemplateId int    `gorm:"column:config_template_id" json:"config_template_id"`
	ImageUrl         string `gorm:"column:image_url" json:"image_url"`
	Envs             string `gorm:"column:envs" json:"envs"`
	BootCmd          string `gorm:"column:boot_cmd" json:"boot_cmd"`
	ImagePullPolicy  string `gorm:"column:image_pull_policy" json:"image_pull_policy"`
}

type MemoryUnitType string

const (
	MemoryUnitMi MemoryUnitType = "Mi"
	MemoryUnitGi MemoryUnitType = "Gi"
)

type TResourceTemplate struct {
	BaseModel
	Name        string         `gorm:"column:name" json:"name"`
	Cpu         float32        `gorm:"column:cpu" json:"cpu"`          // CPU大小
	Memory      int            `gorm:"memory" json:"memory"`           // 内存大小
	MemoryUnit  MemoryUnitType `gorm:"memory_unit" json:"memory_unit"` // 内存单位
	Description string         `gorm:"description" json:"description"`
}

type THealthcheckTemplate struct {
	BaseModel
	Name             string `gorm:"column:name" json:"name"`
	ReadinessEnabled bool   `gorm:"readiness_enabled" json:"readiness_enabled"`
	ReadinessProbe   string `gorm:"column:readiness_probe" json:"readiness_probe"`
	LivenessEnabled  bool   `gorm:"column:liveness_enabled" json:"liveness_enabled"`
	LivenessProbe    string `gorm:"column:liveness_probe" json:"liveness_probe"`
	StartupEnabled   string `gorm:"column:startup_enabled" json:"startup_enabled"`
	StartupProbe     string `gorm:"column:startup_probe" json:"startup_probe"`
	Description      string `gorm:"column:description" json:"description"`
}
