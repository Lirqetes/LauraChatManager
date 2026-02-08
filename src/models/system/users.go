package models

import "lirqetes.ru/thelanc3/laura-server/models"

type User struct {
	models.BaseModel
	TelegramId           int64 `gorm:"column:telegram_id"`
	IsBanned             bool  `gorm:"column:is_banned;default:false"`
	IsReceive            bool  `gorm:"column:is_receive;default:false"`
	IsApprovedStatistics bool  `gorm:"column:is_approved_statistics;default:false"`
}

func (User) TableName() string {
	return "system.users"
}
