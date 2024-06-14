package model

import (
	"cicd/database"
	"fmt"
	"go.uber.org/zap"
)

type TLog struct {
	BaseModel
	Operator string `gorm:"column:operator" json:"operator"`
	Content  string `gorm:"column:content" json:"content"`
}

func (TLog) TableName() string {
	return "t_log"
}
func AddLog(operator, format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	l := &TLog{Operator: operator, Content: content}
	if err := database.DB.Save(l).Error; err != nil {
		zap.L().Error("添加日志异常", zap.Error(err), zap.String("operator", operator), zap.String("content", content))
	}
}
