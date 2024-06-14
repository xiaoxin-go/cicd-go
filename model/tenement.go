package model

type TTenement struct {
	BaseModel
	Name string `gorm:"column:name" json:"name"`
	Desc string `gorm:"column:desc" json:"desc"`
}
