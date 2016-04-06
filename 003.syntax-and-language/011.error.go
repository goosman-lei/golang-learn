package main

import "fmt"
import "errors"

func main() {
	err := UserError{1, errors.New("user error")}
	fmt.Println(err.Error())
}

type UserError struct {
	uid int
	Err error
}

func (e *UserError) Error() string {
	return fmt.Sprintf("user error for uid: %d. errmsg: %s", e.uid, e.Err.Error())
}
