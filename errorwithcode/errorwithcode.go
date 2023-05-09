package errorwithcode

import "fmt"

type ErrorWithCode struct {
	Code int
	Err  error
}

func (e ErrorWithCode) Error() string {
	return fmt.Sprintf("[%v] %s", e.Code, e.Err)
}

func (e ErrorWithCode) IsZero() bool {
	return e.Code == 0 && e.Err == nil
}

func (e ErrorWithCode) IsValid() bool {
	return (e.Code == 0 && e.Err == nil) || (e.Code != 0 && e.Err != nil)
}
