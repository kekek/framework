package error

import "fmt"

type Err struct {
	Code    int
	Message string
}

func New(code int, message string) *Err {
	return &Err{code, message}
}

func (e Err) Error() string {
	return fmt.Sprintf("%s: %d", e.Message, e.Code)
}

func (e Err) String() string {
	return fmt.Sprintf("%s: %d", e.Message, e.Code)
}
