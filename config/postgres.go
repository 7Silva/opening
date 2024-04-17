package config

import (
	"github.com/7Silva/openings/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dsn := "host=localhost user=root password=root dbname=opening port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating schema: %v", err)
		return nil, err
	}

	return db, nil
}
