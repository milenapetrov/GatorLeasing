package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"GatorLeasing/gator-leasing-server/config"
	"GatorLeasing/gator-leasing-server/model"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to database")
		return err
	}

	d.DB = db

	return nil
}

func (d *Database) AutoMigrate() {
	d.DB.AutoMigrate(model.Lease{})
}
