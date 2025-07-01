package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type EmailVerification struct {
	orm.Model
	UserId    uint64 `gorm:"column:user_id"`
	Token     string
	ExpiresAt time.Time `gorm:"column:expires_at"`
	User      *User     `gorm:"foreignKey:UserId"`
	orm.SoftDeletes
}
