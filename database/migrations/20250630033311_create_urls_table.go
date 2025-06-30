package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250630033311CreateUrlsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250630033311CreateUrlsTable) Signature() string {
	return "20250630033311_create_urls_table"
}

// Up Run the migrations.
func (r *M20250630033311CreateUrlsTable) Up() error {
	if !facades.Schema().HasTable("urls") {
		return facades.Schema().Create("urls", func(table schema.Blueprint) {
			table.ID()

			table.UnsignedBigInteger("user_id")
			table.Char("short_code")
			table.Text("original_url")
			table.Boolean("is_active").Default(true)
			table.Char("custom_alias").Nullable()
			table.Char("password_hash").Nullable()
			table.Text("description").Nullable()
			table.UnsignedInteger("clicks_count").Default(0)

			table.TimestampsTz()
			table.SoftDeletesTz()

			table.Foreign("user_id").References("id").On("users").CascadeOnDelete()
			table.Unique("short_code")
			table.Index("is_active")
			table.Index("created_at")
			table.Index("updated_at")
			table.Index("deleted_at")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250630033311CreateUrlsTable) Down() error {
	return facades.Schema().DropIfExists("urls")
}
