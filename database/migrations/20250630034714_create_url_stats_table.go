package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250630034714CreateUrlStatsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250630034714CreateUrlStatsTable) Signature() string {
	return "20250630034714_create_url_stats_table"
}

// Up Run the migrations.
func (r *M20250630034714CreateUrlStatsTable) Up() error {
	if !facades.Schema().HasTable("url_stats") {
		return facades.Schema().Create("url_stats", func(table schema.Blueprint) {
			table.ID()

			table.UnsignedBigInteger("url_id")
			table.Date("date")
			table.UnsignedInteger("clicks_count").Default(0)
			table.UnsignedInteger("unique_click").Default(0)

			table.TimestampsTz()
			table.SoftDeletesTz()

			table.Foreign("url_id").References("id").On("urls").CascadeOnDelete()
			table.Index("date")
			table.Index("clicks_count")
			table.Index("unique_click")
			table.Index("created_at")
			table.Index("updated_at")
			table.Index("deleted_at")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250630034714CreateUrlStatsTable) Down() error {
	return facades.Schema().DropIfExists("url_stats")
}
