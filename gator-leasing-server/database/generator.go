package database

import (
	"math/rand"
	"reflect"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/constants"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/model"
)

func CreateTenants(db *gorm.DB) {
	tenant := model.Tenant{}
	faker.FakeData(tenant)
	db.Create(&tenant)
}

func CreateTenantUsers(db *gorm.DB) {

	for i := 1; i <= 5; i++ {
		tenantUser := model.TenantUser{}
		faker.FakeData(tenantUser)
		tenantUser.ID = uint(i)
		db.Create(&tenantUser)
	}
}

func CreateLeases(db *gorm.DB) {
	for i := 0; i < 1000; i++ {

	}
}

func Generate() {
	_ = faker.AddProvider("tenantIdFaker", func(v reflect.Value) (interface{}, error) {
		return constants.TENANT_ID, nil
	})

	_ = faker.AddProvider("ownerIdFaker", func(v reflect.Value) (interface{}, error) {
		return rand.Intn(5) + 1, nil
	})
}
