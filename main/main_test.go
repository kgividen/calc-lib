package main

import (
	"errors"
	"testing"
)

func TestHandler_WrongNUmberOfArguments(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	if !errors.Is(err, errWrongNUmberofThings) {
		t.Error("wrong error")
	}
}

func TestHandler_InvalidFirstArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"Invalid"})
	if !errors.Is(err, errInvalidArg) {
		t.Error("wrong error")
	}
}
