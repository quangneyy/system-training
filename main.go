package main

import (
	uuid2 "github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type BaseModel struct {
	Id        uuid2.UUID `gorm:"column:id;"`
	Status    string     `gorm:"column:status;"`
	CreatedAt time.Time  `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
}

type Product struct {
	BaseModel
	CategoryId int    `gorm:"column:category_id;"`
	Name       string `gorm:"column:name"`
	//Image       any    `gorm:"column:image"`
	Type        string `gorm:"column:type"`
	Description string `gorm:"column:description"`
}

type ProductUpdate struct {
	Name        *string `gorm:"column:name"`
	CategoryId  *int    `gorm:"column:category_id;"`
	Status      *string `gorm:"column:status;"`
	Type        *string `gorm:"column:type"`
	Description *string `gorm:"column:description"`
}

func (Product) TableName() string { return "products" }

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().UTC()

	newProd := Product{
		BaseModel: BaseModel{
			Status:    "activated",
			CreatedAt: now,
			UpdatedAt: now,
		},
		CategoryId: 1,
		Name:       "Latte",
		//Image:       nil,
		Type:        "drink",
		Description: "",
	}

	if err := db.Table("products").Create(&newProd).Error; err != nil {
		log.Println(err)
	}

	var oldProduct Product

	if err := db.
		Table(Product{}.TableName()).
		Where("id = ", 3).
		First(&oldProduct).Error; err != nil {
		log.Println(err)
	}

	log.Println("Product ID:", oldProduct)

	var prods []Product

	if err := db.
		Table(Product{}.TableName()).
		Limit(10).
		Find(&prods).Error; err != nil {
		log.Println(err)
	}
	log.Println("Products: ", prods)

	//oldProduct.Name = "Capuchino"

	emptyStr := "Latte"

	if err := db.
		Table(Product{}.TableName()).
		Where("id = ?", 3).
		Updates(ProductUpdate{Name: &emptyStr}).Error; err != nil {
		log.Println(err)
	}

	if err := db.
		Table(Product{}.TableName()).
		Where("status not in (?)", []string{"deactivated"}).
		Delete(nil).Error; err != nil {
		log.Println(err)
	}

	uuid, _ := uuid2.NewV7()

	log.Println(uuid.String())
}
