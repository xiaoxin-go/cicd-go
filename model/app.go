package model

type ServiceType string

const (
	ServiceTypeStated    ServiceType = "stated"
	ServiceTypeStateless ServiceType = "stateless"
)

type TApp struct {
	BaseModel
	Name        string      `gorm:"column:name" json:"name"`
	ServiceType ServiceType `gorm:"column:service_type" json:"service_type"`
	Desc        string      `gorm:"desc" json:"desc"`
}
