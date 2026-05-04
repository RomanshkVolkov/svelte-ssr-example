package repository

import (
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/domain"
	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/lg"
	"gorm.io/gorm"
)

func PrintSeedAction(nameTable string, action string) {
	lg.Info("Seeding table: " + nameTable + " " + action + " Success")
}

func AutoMigrateTable(db *gorm.DB, table any) {
	lg.Info("AutoMigrateTable")
	isInitialized := db.Migrator().HasTable(&table)
	if !isInitialized {
		db.AutoMigrate(table)
	}
}

func StartSeeding(db *gorm.DB) {
	lg.Info("StartSeeding")

	AutoMigrateTable(db, &domain.Role{})
	AutoMigrateTable(db, &domain.User{})

	SeedUsers(db)
}
