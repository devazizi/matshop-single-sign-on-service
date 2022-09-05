package rule

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"sso/adapter/db"
)

func EmailMustUnique(db db.DB) validation.RuleFunc {
	return func(value interface{}) error {
		val := value.(string)
		if !db.CheckEmailIsUnique(val) {
			return errors.New("email must unique")
		}

		return nil
	}
}

func CellNumberMustUnique(db db.DB) validation.RuleFunc {
	return func(value interface{}) error {
		val := value.(string)
		if !db.CheckCellNumberIsUnique(val) {
			return errors.New("cell number must unique")
		}

		return nil
	}
}

func NationalCodeMustUnique(db db.DB) validation.RuleFunc {
	return func(value interface{}) error {
		val := value.(string)
		if !db.CheckNationalCodeIsUnique(val) {
			return errors.New("national code must unique")
		}

		return nil
	}
}
