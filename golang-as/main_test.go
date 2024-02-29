package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_Main(t *testing.T) {
	err := errors.New("error")
	aErr := CustomAError{Err: err, Message: err.Error()}
	bErr := CustomBError{Err: aErr, Message: aErr.Error()}

	if diff := cmp.Diff(err, aErr, cmpopts.EquateErrors()); diff != "" {
		t.Errorf("aErr should wrap err: (-got +want)\n%s", diff)
	}

	if diff := cmp.Diff(err, bErr, cmpopts.EquateErrors()); diff != "" {
		t.Errorf("bErr should wrap err: (-got +want)\n%s", diff)
	}

	if diff := cmp.Diff(aErr, bErr, cmpopts.EquateErrors()); diff != "" {
		t.Errorf("bErr should wrap aErr: (-got +want)\n%s", diff)
	}
}
