package database

import (
	"log"

	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
	"github.com/spf13/cobra"
)

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "start seeding",
	RunE:  seeder,
}

func seeder(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}
	admin, _ := libs.HashPassword("admin12345678")
	user, _ := libs.HashPassword("user12345678")

	var data1 = []models.User{
		{Email: "admin@gmail.com", Password: admin, Role: "admin"},
		{Email: "user@gmail.com", Password: user, Role: "user"},
	}

	if res := db.Create(&data1); res.Error != nil {
		return res.Error
	}

	log.Println("Seeding successful")
	return nil
}
