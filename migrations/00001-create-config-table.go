package migrations

import "github.com/jinzhu/gorm"

func CREATE_CONFIG_TABLE_AND_ACL() MigrationStep {
	type AccessControl struct {
		gorm.Model
		AccessID string
		UserID   string
		Read     bool
		Write    bool
		Delete   bool
	}

	type Config struct {
		gorm.Model
		Name     string
		Value    string
		AccessID string
	}
	return MigrationStep{
		Name: "00001-create-config-table",
		Down: func(tx *gorm.DB) error {
			return tx.DropTableIfExists(Config{}, AccessControl{}).Error
		},
		Up: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(Config{}, AccessControl{}).Error; err != nil {
				return err
			}
			if err := tx.Create(
				&Config{
					Name:     "echo",
					Value:    "I am alive",
					AccessID: "config:echo",
				},
			).Error; err != nil {
				return err
			}
			if err := tx.Create(
				&AccessControl{
					AccessID: "config:echo",
					UserID:   "_system",
					Read:     true,
					Write:    true,
				},
			).Error; err != nil {
				return err
			}
			return nil
		},
	}
}
