package AppContact

type ErrorStruct struct {
	Error string
}

const (
	ErrorUserContactExists = "User contact already exists"
)

func ErrorUserAlreadyExists() []ErrorStruct {
	return []ErrorStruct{{Error: ErrorUserContactExists}}
}
