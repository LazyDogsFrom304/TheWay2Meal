package service

type capError struct {
	e   error
	msg string
}

func (ce capError) Error() string {
	return ce.e.Error() + ce.msg
}
