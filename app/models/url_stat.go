package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type UrlStat struct {
	orm.Model
	UrlID       uint64
	Date        time.Time
	ClicksCount uint `gorm:"column:clicks_count"`
	UniqueClick uint `gorm:"column:unique_click"`
	Url         Url  `gorm:"foreignKey:UrlID"`
	orm.SoftDeletes
}
