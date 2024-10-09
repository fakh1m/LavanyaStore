package fakers

import (
	"log"
	"math"
	"math/rand/v2"
	"time"

	"github.com/fakh1m/LavanyaStore/app/models"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *models.Product {
	user := UserFaker(db)
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	name := faker.Name()
	return &models.Product{
		ID:               uuid.New().String(),
		User:             models.User{},
		UserID:           user.ID,
		Sku:              slug.Make(name),
		Name:             name,
		Slug:             slug.Make(name),
		Price:            decimal.NewFromFloat(fakerPrice()),
		Stock:            rand.IntN(100),
		Weight:           decimal.NewFromFloat(rand.Float64()),
		ShortDescription: faker.Paragraph(),
		Description:      faker.Paragraph(),
		Status:           1,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}

func fakerPrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.IntN(8)), rand.IntN(2)+1)
}

func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val*div)) / div
}
