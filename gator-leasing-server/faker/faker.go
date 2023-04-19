package faker

import (
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/constants"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
	"github.com/shopspring/decimal"
	"github.com/tjarratt/babble"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func InitializeFaker() {
	amenityOptions := strings.Split(constants.AMENITIES, ",")
	numAmenities, _ := faker.RandomInt(0, len(amenityOptions), 1)
	amenities, _ := faker.RandomInt(0, len(amenityOptions)-1, numAmenities[0])
	println(amenities)

	_ = faker.AddProvider("nameFaker", func(v reflect.Value) (interface{}, error) {
		babbler := babble.NewBabbler()
		babbler.Count = 1
		return cases.Title(language.English).String(babbler.Babble()), nil
	})

	_ = faker.AddProvider("streetFaker", func(v reflect.Value) (interface{}, error) {
		houseNumber, _ := faker.RandomInt(1, 99999, 1)
		babbler := babble.NewBabbler()
		name := cases.Title(language.English).String(babbler.Babble())
		suffixes := strings.Split(constants.STREET_SUFFIXES, ",")
		suffix, _ := faker.RandomInt(0, len(suffixes)-1, 1)

		return strconv.Itoa(houseNumber[0]) + " " + name + " " + suffixes[suffix[0]], nil
	})

	_ = faker.AddProvider("cityFaker", func(v reflect.Value) (interface{}, error) {
		babbler := babble.NewBabbler()
		babbler.Count = 1
		return cases.Title(language.English).String(babbler.Babble()), nil
	})

	_ = faker.AddProvider("stateFaker", func(v reflect.Value) (interface{}, error) {
		states := strings.Split(constants.STATES, ",")
		state, _ := faker.RandomInt(0, len(states)-1, 1)
		return states[state[0]], nil
	})

	_ = faker.AddProvider("zipCodeFaker", func(v reflect.Value) (interface{}, error) {
		zipCode, _ := faker.RandomInt(10000, 99999, 1)
		return strconv.Itoa(zipCode[0]), nil
	})

	_ = faker.AddProvider("startDateFaker", func(v reflect.Value) (interface{}, error) {
		randomYear, _ := faker.RandomInt(-2, 0, 1)
		randomMonth, _ := faker.RandomInt(-11, 0, 1)
		randomDay, _ := faker.RandomInt(-30, 0, 1)
		return time.Now().AddDate(randomYear[0], randomMonth[0], randomDay[0]), nil
	})

	_ = faker.AddProvider("endDateFaker", func(v reflect.Value) (interface{}, error) {
		randomYear, _ := faker.RandomInt(0, 2, 1)
		randomMonth, _ := faker.RandomInt(0, 11, 1)
		randomDay, _ := faker.RandomInt(0, 30, 1)
		return time.Now().AddDate(randomYear[0], randomMonth[0], randomDay[0]), nil
	})

	_ = faker.AddProvider("rentFaker", func(v reflect.Value) (interface{}, error) {
		return decimal.NewFromFloat(0.01 + rand.Float64()*(10000-0.01)).Round(2), nil
	})

	_ = faker.AddProvider("utilitiesFaker", func(v reflect.Value) (interface{}, error) {
		return decimal.NewFromFloat(0.01 + rand.Float64()*(300-0.01)).Round(2), nil
	})

	_ = faker.AddProvider("parkingCostFaker", func(v reflect.Value) (interface{}, error) {
		return decimal.NewFromFloat(0.01 + rand.Float64()*(300-0.01)).Round(2), nil
	})

	_ = faker.AddProvider("parkingFaker", func(v reflect.Value) (interface{}, error) {
		return true, nil
	})

	_ = faker.AddProvider("bathsFaker", func(v reflect.Value) (interface{}, error) {
		baths, _ := faker.RandomInt(1, 16, 1)
		return decimal.NewFromFloat(float64(baths[0]) / 2).Round(1), nil
	})

	_ = faker.AddProvider("amenitiesFaker", func(v reflect.Value) (interface{}, error) {
		amenityOptions := strings.Split(constants.AMENITIES, ",")
		numAmenities, _ := faker.RandomInt(0, len(amenityOptions), 1)
		amenities, _ := faker.RandomInt(0, len(amenityOptions)-1, numAmenities[0])
		amenitiesAsString := ""
		for _, a := range amenities {
			amenitiesAsString += amenityOptions[a] + ","
		}
		if amenitiesAsString == "" {
			return "", nil
		}
		return amenitiesAsString[:len(amenitiesAsString)-1], nil
	})

	_ = faker.AddProvider("appliancesFaker", func(v reflect.Value) (interface{}, error) {
		applianceOptions := strings.Split(constants.APPLIANCES, ",")
		numAppliances, _ := faker.RandomInt(0, len(applianceOptions), 1)
		appliances, _ := faker.RandomInt(0, len(applianceOptions)-1, numAppliances[0])
		appliancesAsString := ""
		for _, a := range appliances {
			appliancesAsString += applianceOptions[a] + ","
		}
		if appliancesAsString == "" {
			return "", nil
		}
		return appliancesAsString[:len(appliancesAsString)-1], nil
	})

	_ = faker.AddProvider("descriptionFaker", func(v reflect.Value) (interface{}, error) {
		babbler := babble.NewBabbler()
		babbler.Count = 10
		babbler.Separator = " "
		return babbler.Babble(), nil
	})

	_ = faker.AddProvider("sortDirectionFaker", func(v reflect.Value) (interface{}, error) {
		sortDirection, _ := faker.RandomInt(0, 1, 1)
		return enums.SortDirection(sortDirection[0]), nil
	})

	_ = faker.AddProvider("filtersFaker", func(v reflect.Value) (interface{}, error) {
		babbler := babble.NewBabbler()
		babbler.Count = 1
		return babbler.Babble() + " " + babbler.Babble() + " '" + babbler.Babble() + "'", nil
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
