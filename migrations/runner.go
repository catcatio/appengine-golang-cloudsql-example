package migrations

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type MigrationStep struct {
	Name string
	Up   func(tx *gorm.DB) error
	Down func(tx *gorm.DB) error
}

var Migrations = []MigrationStep{
	CREATE_CONFIG_TABLE_AND_ACL(),
}

type MigrationLog struct {
	*gorm.Model
	Name string `gorm:"unique"`
}

func Migrate(db *gorm.DB) {
	fmt.Printf("\n\n🏗 Start Database versioning migration....\n")
	db.AutoMigrate(&MigrationLog{})
	for _, migrate := range Migrations {
		var n int
		db.Model(&MigrationLog{}).Where(&MigrationLog{
			Name: migrate.Name,
		}).Count(&n)

		if n > 0 {
			fmt.Printf("⏭ %s Skipped \n", migrate.Name)
			continue
		}
		tx := db.Debug().Begin()
		if err := migrate.Up(tx); err != nil {
			fmt.Printf("💥 %s Failed\n", migrate.Name)
			panic(err)
		}
		if err := tx.Commit().GetErrors(); len(err) > 0 {
			tx.Rollback()
			panic(err)
		} else {
			db.Create(&MigrationLog{
				Name: migrate.Name,
			})
			fmt.Printf("✅ %s Success\n", migrate.Name)
		}
	}
	fmt.Printf("============\n🎉🚀Oh yeahhhh, Database migration completed !!\n\n")
}
