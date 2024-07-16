package model

type ServiceType string

const (
	ServiceTypeStated    ServiceType = "stated"    // 有状态应用
	ServiceTypeStateless ServiceType = "stateless" // 无状态应用
)

type TApp struct {
	BaseModel
	Name        string      `gorm:"column:name" json:"name"`
	ServiceType ServiceType `gorm:"column:service_type" json:"service_type"`
	Description string      `gorm:"description" json:"description"`
}

func (*TApp) TableName() string {
	return "t_app"
}
