package models

import (
	"github.com/goravel/framework/database/orm"
)

type Url struct {
	orm.Model
	UserId       *uint64
	ShortCode    string  `gorm:"column:short_code"`
	OriginalUrl  string  `gorm:"column:original_url"`
	IsActive     bool    `gorm:"column:is_active"`
	CustomAlias  *string `gorm:"column:custom_alias"`
	PasswordHash *string `gorm:"column:password_hash"`
	Description  *string
	ClicksCount  uint      `gorm:"column:clicks_count"`
	User         User      `gorm:"foreignKey:UserId"`
	Clicks       []Click   `gorm:"foreignKey:UrlID"`
	UrlStats     []UrlStat `gorm:"foreignKey:UrlID"`
	orm.SoftDeletes
}
