package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250630061740CreateEmailVerificationsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250630061740CreateEmailVerificationsTable) Signature() string {
	return "20250630061740_create_email_verifications_table"
}

// Up Run the migrations.
func (r *M20250630061740CreateEmailVerificationsTable) Up() error {
	if !facades.Schema().HasTable("email_verifications") {
		return facades.Schema().Create("email_verifications", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("user_id")
			table.String("token")
			table.Timestamp("expires_at")
			table.TimestampsTz()
			table.SoftDeletesTz()

			table.Foreign("user_id").References("id").On("users").CascadeOnDelete()
			table.Index("token")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250630061740CreateEmailVerificationsTable) Down() error {
	return facades.Schema().DropIfExists("email_verifications")
}
