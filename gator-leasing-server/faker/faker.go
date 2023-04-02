package faker

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/shopspring/decimal"
)

func InitializeFaker() {
	_ = faker.AddProvider("ownerIdFaker", func(v reflect.Value) (interface{}, error) {
		return rand.Intn(5) + 1, nil
	})

	_ = faker.AddProvider("contactsEntityFaker", func(v reflect.Value) (interface{}, error) {
		fakeContacts := []entity.Contact{}
		for i := 0; i < 3; i++ {
			fakeContact := entity.Contact{}
			faker.FakeData(&fakeContact)
			fakeContacts = append(fakeContacts, fakeContact)
		}
		return fakeContacts, nil
	})

	_ = faker.AddProvider("createLeaseStartDateFaker", func(v reflect.Value) (interface{}, error) {
		randomYear, _ := faker.RandomInt(-3, 1, 1)
		randomMonth, _ := faker.RandomInt(-12, 1, 1)
		randomDay, _ := faker.RandomInt(-31, 1, 1)
		return time.Now().AddDate(randomYear[0], randomMonth[0], randomDay[0]), nil
	})

	_ = faker.AddProvider("createLeaseEndDateFaker", func(v reflect.Value) (interface{}, error) {
		randomYear, _ := faker.RandomInt(-1, 3, 1)
		randomMonth, _ := faker.RandomInt(-1, 12, 1)
		randomDay, _ := faker.RandomInt(-1, 31, 1)
		return time.Now().AddDate(randomYear[0], randomMonth[0], randomDay[0]), nil
	})

	_ = faker.AddProvider("createLeaseRentFaker", func(v reflect.Value) (interface{}, error) {
		return decimal.NewFromFloat(0.01 + rand.Float64()*(10000-0.01)).Round(2), nil
	})
}

func FakeData(x interface{}) error {
	return faker.FakeData(x)
}

func FakeMany[T any](x *T, n int) []*T {
	slice := []*T{}
	for i := 0; i < n; i++ {
		fake := new(T)
		FakeData(fake)
		slice = append(slice, fake)
	}

	return slice
}
