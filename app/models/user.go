package models

import (
	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
	"github.com/linkeunid/api.linkeun.com/database/factories"
)

type User struct {
	orm.Model
	Name     string
	Email    string
	Password string
	orm.SoftDeletes
}

func (u *User) Factory() factory.Factory {
	return &factories.UserFactory{}
}
