package errors

type BadRequestError struct {
	Msg string `json:"error"`
}

func (e *BadRequestError) Error() string {
	return e.Msg
}

func (e *BadRequestError) Is(target error) bool {
	_, ok := target.(*BadRequestError)

	return ok
}
