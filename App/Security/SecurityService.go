package AppSecurity

import (
	SecurityCommand "gontact/App/Security/Command"
	AppSecuritySecurityDTO "gontact/App/Security/DTO"
	DomainUsers "gontact/Domain/Users"
	DomainValidator "gontact/Domain/Validator"
	DomainValidatorVO "gontact/Domain/Validator/VO"
)

func GetCurrentUser(command SecurityCommand.SecurityCommand) AppSecuritySecurityDTO.TokenUserDTO {
	vo := DomainUsers.GetUserTokenVO(command.GetToken())
	entity := DomainUsers.GetUserByToken(vo)

	return GetTokenUserDTO(entity)
}
func Validate(i interface{}) []*DomainValidatorVO.ValidatorVO {
	return DomainValidator.Validate(i)
}
