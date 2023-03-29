package shared

type BadRequestError struct {
	Msg string `json:"error"`
}

func (e *BadRequestError) Error() string {
	return e.Msg
}
