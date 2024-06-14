package model

type TNamespace struct {
	BaseModel
	TenementId int    `gorm:"column:tenement_id" json:"tenement_id"`
	Name       string `gorm:"column:name" json:"name"`
	Uid        string `gorm:"column:uid" json:"uid"`
	Status     string `gorm:"column:status" json:"status"`
	Labels     string `gorm:"column:labels" json:"labels"`
}
