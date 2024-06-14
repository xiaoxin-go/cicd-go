package model

type TUser struct {
	BaseModel
	Username string `gorm:"column:username" json:"username"`
	Enabled  uint32 `gorm:"column:enabled" json:"enabled"`
}
