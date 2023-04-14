package validator

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/errors"
)

type Validator struct {
	validator  *validator.Validate
	translator ut.Translator
}

func New() *Validator {
	en := en.New()
	translator, ok := ut.New(en, en).GetTranslator("en")
	if !ok {
		log.Fatal("could not find translator")
	}
	validator := &Validator{validator: validator.New(), translator: translator}
	validator.registerCustomValidations()
	validator.registerTranslations()
	return validator
}

func (v *Validator) registerCustomValidations() {
	v.validator.RegisterCustomTypeFunc(decimalCustomTypeFunc, decimal.Decimal{})
	v.validator.RegisterValidation("dmin", decimalMinValidator)
	v.validator.RegisterValidation("dmax", decimalMaxValidator)
}

func (v *Validator) registerTranslations() {
	v.validator.RegisterTranslation("required", v.translator, registerRequiredTranslation, translateRequired)
	v.validator.RegisterTranslation("min", v.translator, registerMinTranslation, translateMin)
	v.validator.RegisterTranslation("dmin", v.translator, registerDminTranslation, translateDmin)
	v.validator.RegisterTranslation("gtfield", v.translator, registerGtfieldTranslation, translateGtfield)
}

func (v *Validator) Struct(x interface{}) []error {
	errs := []error{}
	err := v.validator.Struct(x)
	if err == nil {
		return nil
	}
	if _, invalid := err.(*validator.InvalidValidationError); invalid {
		return append(errs, &errors.InternalServerError{Msg: err.Error()})
	}
	validationErrs := err.(validator.ValidationErrors)
	for _, e := range validationErrs {
		errs = append(errs, &errors.BadRequestError{Msg: e.Translate(v.translator)})
	}
	return errs
}

// custom type funcs
func decimalCustomTypeFunc(field reflect.Value) interface{} {
	if value, ok := field.Interface().(decimal.Decimal); ok {
		if value.Equal(decimal.Zero) {
			return ""
		}
		return value.StringFixed(2)
	}
	return nil
}

// validators
func decimalMinValidator(fl validator.FieldLevel) bool {
	min, _ := decimal.NewFromString(fl.Param())
	val, _ := decimal.NewFromString(fl.Field().String())
	return val.GreaterThanOrEqual(min)
}

func decimalMaxValidator(fl validator.FieldLevel) bool {
	max, _ := decimal.NewFromString(fl.Param())
	val, _ := decimal.NewFromString(fl.Field().String())
	return val.LessThanOrEqual(max)
}

// register translation funcs
func registerRequiredTranslation(ut ut.Translator) error {
	return ut.Add("required", "{0} is required", true)
}

func registerMinTranslation(ut ut.Translator) error {
	return ut.Add("min", "{0} ({1}) must be at least {2}", true)
}

func registerDminTranslation(ut ut.Translator) error {
	return ut.Add("dmin", "{0} ({1}) must be at least {2}", true)
}

func registerGtfieldTranslation(ut ut.Translator) error {
	return ut.Add("gtfield", "{0} ({1}) must be greater than {2}", true)
}

// translation funcs
func translateRequired(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("required", fe.StructField())
	return t
}

func translateMin(ut ut.Translator, fe validator.FieldError) string {
	t := ""
	switch fe.Value().(type) {
	case string:
		t, _ = ut.T("min", fe.StructField(), strconv.Itoa(len((fe.Value().(string)))), fe.Param()+" characters")
	default:
		t, _ = ut.T("min", fe.StructField(), fmt.Sprintf("%v", fe.Value()), fe.Param())
	}
	return t
}

func translateDmin(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("dmin", fe.StructField(), fe.Value().(string), fe.Param())
	return t
}

func translateGtfield(ut ut.Translator, fe validator.FieldError) string {
	t := ""
	switch fe.Value().(type) {
	case time.Time:
		t, _ = ut.T("gtfield", fe.StructField(), fe.Value().(time.Time).Format("1/2/2006"), fe.Param())
	default:
		t, _ = ut.T("gtfield", fe.StructField(), fmt.Sprintf("%v", fe.Value()), fe.Param())
	}
	return t
}
