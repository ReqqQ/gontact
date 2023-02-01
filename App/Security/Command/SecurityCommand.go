package AppSecuritySecurityCommand

type SecurityCommand struct {
	Token string `reqHeader:"X-Cont-Token" validate:"required"`
}

func (r SecurityCommand) GetToken() string {
	return r.Token
}
