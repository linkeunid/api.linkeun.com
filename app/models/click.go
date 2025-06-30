package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Click struct {
	orm.Model
	UrlID     uint64
	ClickedAt time.Time `gorm:"column:clicked_at"`
	IpAddress *string   `gorm:"column:ip_address"`
	UserAgent *string   `gorm:"column:user_agent"`
	Referrer  *string
	Browser   *string
	Url       Url `gorm:"foreignKey:UrlID"`
	orm.SoftDeletes
}
