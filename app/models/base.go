package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model interface {}

type BaseModel struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

/*
|--------------------------------------------------------------------------
| initialize create time
|--------------------------------------------------------------------------
*/
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

/*
|--------------------------------------------------------------------------
| initialize update time
|--------------------------------------------------------------------------
*/
func (m *BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	now := time.Now()
	now = now.Round(time.Second)
	scope.SetColumn("UpdatedAt", now)
	return nil
}
