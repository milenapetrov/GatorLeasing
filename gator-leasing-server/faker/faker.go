package faker

import (
	"math/rand"
	"reflect"

	"github.com/bxcodec/faker/v4"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
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
}

func FakeData(x interface{}) {
	err := faker.FakeData(x)
	println(err)
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
