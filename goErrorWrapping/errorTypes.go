package main

type appError struct {
	msg   string
	fault error
}

func (e appError) Unwrap() error {
	return e.fault
}

func (e appError) Error() string {
	return e.msg
}

type laptopError struct {
	msg   string
	fault error
}

func (e laptopError) Unwrap() error {
	return e.fault
}

func (e laptopError) Error() string {
	return e.msg
}

type programError struct {
	msg   string
	fault error
}

func (e programError) Unwrap() error {
	return e.fault
}

func (e programError) Error() string {
	return e.msg
}

type locateError struct {
	msg   string
	fault error
}

func (e locateError) Unwrap() error {
	return e.fault
}

func (e locateError) Error() string {
	return e.msg
}

type goError struct {
	msg   string
	fault error
}

func (e goError) Unwrap() error {
	return e.fault
}

func (e goError) Error() string {
	return e.msg
}

type buyError struct {
	msg   string
	fault error
}

func (e buyError) Unwrap() error {
	return e.fault
}

func (e buyError) Error() string {
	return e.msg
}
