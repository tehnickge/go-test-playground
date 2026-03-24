package models

import (
	"time"

	"gorm.io/gorm"
)

// User — основная модель пользователя
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UUID      string         `gorm:"type:uuid;uniqueIndex;not null" json:"uuid"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"`       // никогда не отдаём пароль в JSON
	Role      string         `gorm:"size:20;default:user" json:"role"` // user, admin, moderator
	Status    string         `gorm:"size:20;default:active" json:"status"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // поддержка мягкого удаления (Soft Delete)

	// Дополнительные полезные поля
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	LastLoginAt     *time.Time `json:"last_login_at,omitempty"`
}

// TableName — можно задать своё имя таблицы (по умолчанию будет "users")
func (User) TableName() string {
	return "users"
}
