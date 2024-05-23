package initializers

import "github.com/ApesJs/test-be/models"

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.Post{},
	)
	if err != nil {
		panic(err)
	}
}
