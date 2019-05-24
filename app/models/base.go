package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model interface {
	GetId() int
}

type BaseModel struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *BaseModel) GetId() int {
	return m.Id
}

func (m *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	now := time.Now()
	now = now.Round(time.Second)
	if m.CreatedAt.IsZero() {
		scope.SetColumn("CreatedAt", now)
	}
	if m.UpdatedAt.IsZero() {
		scope.SetColumn("UpdatedAt", now)
	}
	return nil
}

func (m *BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	now := time.Now()
	now = now.Round(time.Second)
	scope.SetColumn("UpdatedAt", now)
	return nil
}

func (m *BaseModel) RevisionKey() int {
	return m.Id
}
