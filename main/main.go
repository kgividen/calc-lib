package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/kgividen/calc-lib"
)

func main() {
	handler := NewHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

}

type Handler struct {
	stdout     io.Writer
	calculator *calc.Addition
}

func NewHandler(stdout io.Writer, calculator *calc.Addition) *Handler {
	return &Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongNUmberofThings
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	b, err := strconv.Atoi(args[1])

	if err != nil {
		return err
	}

	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)

	if err != nil {
		return err
	}

	return nil
}

var errWrongNUmberofThings = errors.New("Wrong number of things.")
var errInvalidArg = errors.New("Invalid Arg.")
