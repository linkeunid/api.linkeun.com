package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"github.com/linkeunid/api.linkeun.com/database/migrations"
	"github.com/linkeunid/api.linkeun.com/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20240915060148CreateUsersTable{},
		&migrations.M20250630033311CreateUrlsTable{},
		&migrations.M20250630033316CreateClicksTable{},
		&migrations.M20250630034714CreateUrlStatsTable{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.DatabaseSeeder{},
		&seeders.UserSeeder{},
	}
}
