package seeders

import (
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/models"

	"github.com/go-faker/faker/v4"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	randPass := faker.Word()
	facades.Log().Debug("Your password is: ", randPass)
	secret, err := facades.Crypt().EncryptString(randPass)
	if err != nil {
		return err
	}
	user := models.User{
		Name:     "Hanivan Rizky S",
		Email:    "hanivan@linkeun.com",
		Password: secret,
	}

	return facades.Orm().Query().Create(&user)
}
