package seeder

import (
	"service_social_media/app/databases/faker"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.UserFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error

		if err != nil {
			return err
		}
	}

	return nil
}
