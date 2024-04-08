package functions

import (
	"github.com/7Silva/openings/config"
	"github.com/7Silva/openings/schemas"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitDB() {
	logger = config.GetLogger("db")
	db = config.GetSQLite()
}

func GetOpening(id string) (schemas.Opening, error) {
	opening := schemas.Opening{}
	err := db.First(&opening, id).Error
	return opening, err
}

func GetOpenings() ([]schemas.Opening, error) {
	openings := []schemas.Opening{}
	err := db.Find(&openings).Error
	return openings, err
}

func OpeningExists(id string) bool {
	return db.First(&schemas.Opening{}, id).Error == nil
}

func DeleteOpening(id string) bool {
	return db.Delete(&schemas.Opening{}, id).Error == nil
}

func CreateOpening(opening schemas.Opening) (schemas.Opening, error) {
	err := db.Create(&opening).Error

	if err != nil {
		logger.Error(err.Error())
	}

	return opening, err
}

func UpdateOpening(opening schemas.Opening) (schemas.Opening, error) {
	err := db.Model(&opening).Updates(opening).Error
	if err != nil {
		logger.Error(err.Error())
		return opening, err
	}

	return opening, nil
}
