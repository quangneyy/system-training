package common

import (
	uuid2 "github.com/google/uuid"
	"time"
)

type BaseModel struct {
	Id        uuid2.UUID `gorm:"column:id;" json:"id"`
	Status    string     `gorm:"column:status;" json:"status"`
	CreatedAt time.Time  `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;" json:"updated_at"`
}

func GenNewModel() BaseModel {
	now := time.Now().UTC()
	newId, _ := uuid2.NewV7()

	return BaseModel{
		Id:        newId,
		Status:    "activated",
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func GenUUID() uuid2.UUID {
	newID, _ := uuid2.NewV7()
	return newID
}

func ParseUUID(s string) uuid2.UUID {
	return uuid2.MustParse(s)
}
