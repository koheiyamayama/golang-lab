package main

import "errors"

type CustomAError struct {
	Err     error
	Message string
}

type CustomBError struct {
	Err     error
	Message string
}

func (c CustomAError) Error() string {
	return "CustomAError: " + c.Message
}

func (c CustomAError) Unwrap() error {
	return c.Err
}

func (c CustomBError) Error() string {
	return "CustomBError: " + c.Message
}

func (c CustomBError) Unwrap() error {
	return c.Err
}

func main() {
	err := errors.New("error")
	aErr := CustomAError{Err: err, Message: err.Error()}
	bErr := CustomBError{Err: aErr, Message: aErr.Error()}

	if errors.Is(aErr, err) {
		println("aErr wraps err")
	}

	if errors.Is(bErr, err) {
		println("bErr wraps err")
	}

	if errors.Is(bErr, aErr) {
		println("bErr wraps aErr")
	} else if errors.Is(aErr, bErr) {
		println("aErr wraps bErr")
	}
}
