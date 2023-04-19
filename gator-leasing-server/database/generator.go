package database

import (
	"bufio"
	stdErrors "errors"
	"log"
	"os"
	"strings"

	random "github.com/bxcodec/faker/v4"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/faker"
)

func (d *Database) Generate() {
	faker.InitializeFaker()

	tenantId := d.createTenant()
	tenantUserIds := d.createTenantUsers(tenantId)
	d.createLeases(tenantUserIds)
}

func (d *Database) createTenant() uint {
	tenant := &dto.Tenant{
		ID:   1,
		Name: "Default Tenant",
	}
	d.DB.Create(tenant)

	return tenant.ID
}

func (d *Database) createTenantUsers(tenantId uint) []uint {
	tenantUserIds := []uint{}
	file, err := os.Open("users/user_info.txt")
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			userInfo := strings.Split(scanner.Text(), ",")
			if len(userInfo) < 3 {
				log.Fatal("invalid user info")
			}
			tenantUser := &dto.TenantUser{
				UserID:      userInfo[0],
				TenantID:    tenantId,
				InvitedAs:   enums.Member,
				Email:       userInfo[1],
				PhoneNumber: userInfo[2],
			}
			d.DB.Create(tenantUser)
			tenantUserIds = append(tenantUserIds, tenantUser.ID)
		}
	} else if !stdErrors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	if len(tenantUserIds) == 0 {
		for i := 0; i < 5; i++ {
			tenantUser := &dto.TenantUser{
				TenantID:  tenantId,
				InvitedAs: enums.Member,
			}
			faker.FakeData(tenantUser)
			d.DB.Create(tenantUser)
			tenantUserIds = append(tenantUserIds, tenantUser.ID)
		}
	}

	return tenantUserIds
}

func (d *Database) createLeases(ownerIds []uint) {
	for i := 0; i < 1000; i++ {
		ownerId, _ := random.RandomInt(0, len(ownerIds)-1, 1)
		lease := &dto.Lease{
			OwnerID: ownerIds[ownerId[0]],
		}
		faker.FakeData(lease)
		d.DB.Create(lease)
	}
}
