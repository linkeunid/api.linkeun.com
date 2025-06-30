package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250630033316CreateClicksTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250630033316CreateClicksTable) Signature() string {
	return "20250630033316_create_clicks_table"
}

// Up Run the migrations.
func (r *M20250630033316CreateClicksTable) Up() error {
	if !facades.Schema().HasTable("clicks") {
		return facades.Schema().Create("clicks", func(table schema.Blueprint) {
			table.ID()

			table.UnsignedBigInteger("url_id")
			table.TimestampTz("clicked_at")
			table.Char("ip_address").Nullable()
			table.Text("user_agent").Nullable()
			table.Text("referrer").Nullable()
			table.Char("browser").Nullable()

			table.TimestampsTz()
			table.SoftDeletesTz()

			table.Foreign("url_id").References("id").On("urls")
			table.Index("clicked_at")
			table.Index("created_at")
			table.Index("updated_at")
			table.Index("deleted_at")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250630033316CreateClicksTable) Down() error {
	return facades.Schema().DropIfExists("clicks")
}
