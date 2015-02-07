package parser

import (
	"github.com/sorah/logmetro/core"
	"strings"
)

type Ltsv struct {
	Parser
}

func NewLtsvParser() (*Ltsv, error) {
	return &Ltsv{}, nil
}

func (parser *Ltsv) Parse(str string) (*core.LogData, error) {
	data := make(map[string]string, 10)
	for _, entry := range strings.Split(str, "\t") {
		keyval := strings.SplitN(entry, ":", 2)
		if len(keyval) > 1 {
			data[keyval[0]] = keyval[1]
		} else {
			data[keyval[0]] = ""
		}
	}
	logdata := &core.LogData{Data: data}
	return logdata, nil
}
