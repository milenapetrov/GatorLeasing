package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/config"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/faker"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) GetConnection(dbConfig *config.DBConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Address,
		dbConfig.Name,
		dbConfig.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("could not connect to database")
		return err
	}

	d.DB = db

	return nil
}

func (d *Database) AutoMigrate() {
	err := d.DB.AutoMigrate(&dto.Lease{}, &dto.Address{}, &dto.Contact{}, &dto.Tenant{}, &dto.TenantUser{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (d *Database) Generate() {
	faker.InitializeFaker()
}
