package database

import (
	"github.com/bxcodec/faker/v4"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/constants"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
)

func (d *Database) createTenants() {
	tenant := dto.Tenant{}
	tenant.ID = constants.TENANT_ID
	faker.FakeData(tenant)
	d.DB.Create(&tenant)
}

func (d *Database) createTenantUsers() {
	tenantUser := &dto.TenantUser{}
	faker.FakeData(tenantUser)

	for i := 1; i <= 5; i++ {
		tenantUser = &dto.TenantUser{}
		faker.FakeData(tenantUser)
		tenantUser.ID = uint(i)
		d.DB.Create(&tenantUser)
	}
}

func (d *Database) createLeases() {
	for i := 0; i < 1000; i++ {

	}
}
