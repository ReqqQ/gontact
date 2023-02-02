package Response

import AppResponseCommand "gontact/App/Response/Command"

const UserNotExists = "User not exists.Check again"

func Error(i interface{}) AppResponseCommand.ErrorResponseStructure {
	return AppResponseCommand.ErrorResponseStructure{
		Errors: i,
	}
}
func Success(i interface{}) AppResponseCommand.SuccessResponseStructure {
	return AppResponseCommand.SuccessResponseStructure{
		Data: i,
	}
}
