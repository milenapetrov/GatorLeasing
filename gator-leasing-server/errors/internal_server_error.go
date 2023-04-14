package errors

type InternalServerError struct {
	Msg string `json:"error"`
}

func (e *InternalServerError) Error() string {
	return e.Msg
}

func (e *InternalServerError) Is(target error) bool {
	_, ok := target.(*InternalServerError)

	return ok
}
