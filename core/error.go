package core

import (
	"fmt"
)

const (
	InvalidConfigurationError = 1450
	UnknownInputTypeError     = 1451
	UnknownParserTypeError    = 1452

	ParserError = 1460
)

var errorTypes = map[int]string{
	InvalidConfigurationError: "Invalid configuration",
	UnknownInputTypeError:     "Unknown input type",
	UnknownParserTypeError:    "Unknown parser type",
	ParserError:               "Parser error",
}

type Error struct {
	error
	Code    int
	Kind    string
	Message string
}

func NewError(code int, message string) *Error {
	kind, ok := errorTypes[code]
	if !ok {
		panic(fmt.Errorf("Unknown error code %d", code))
	}
	return &Error{
		Code:    code,
		Kind:    kind,
		Message: message,
	}
}

func (e *Error) Error() string {
	if e.Message == "" {
		return e.Kind
	} else {
		return e.Kind + ": " + e.Message
	}
}
