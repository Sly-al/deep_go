package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errs []error
}

func (e *MultiError) Error() string {
	res := "1 error occured"
	if len(e.errs) > 1 {
		res = fmt.Sprintf("%d errors occured:\n", len(e.errs))
	}
	for i := range e.errs {
		res += fmt.Sprintf("\t* %s", e.errs[i])
	}

	return res + "\n"
}

func Append(err error, errs ...error) *MultiError {
	t, ok := err.(*MultiError)
	if !ok {
		t = &MultiError{}
	}

	for i := range errs {
		t.errs = append(t.errs, errs[i])
	}

	return t
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
