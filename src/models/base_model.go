package models

import "time"

type BaseModel struct {
	Id        int64     `gorm:"column:id;autoIncrement;primaryKey;unique"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
