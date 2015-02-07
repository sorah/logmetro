package parser

import (
	"github.com/sorah/logmetro/core"
)

type Parser interface {
	Parse(str string) (*core.LogData, error)
}

func NewParser(typeName string) (Parser, error) {
	switch typeName {
	case "ltsv":
		return NewLtsvParser()
	default:
		return nil, core.NewError(core.UnknownParserTypeError, "")
	}
}
