package DomainValidator

import (
	"github.com/go-playground/validator/v10"
	DomainValidatorVO "gontact/Domain/Validator/VO"
)

func Validate(i interface{}) []*DomainValidatorVO.ValidatorVO {
	var validate = validator.New()
	var errors []*DomainValidatorVO.ValidatorVO
	err := validate.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element DomainValidatorVO.ValidatorVO
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
