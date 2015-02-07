package parser

import (
	"testing"
)

func TestLtsvParse(t *testing.T) {
	parser, err := NewLtsvParser()
	if err != nil {
		t.Error(err)
	}

	data, err := parser.Parse("a:foo\tb:bar\tc:1:2:3")

	val := data.Data.(map[string]string)["a"]
	if val != "foo" {
		t.Errorf("a, expected:foo actual:%v", val)
	}
	val = data.Data.(map[string]string)["b"]
	if val != "bar" {
		t.Errorf("b, expected:bar actual:%v", val)
	}
	val = data.Data.(map[string]string)["c"]
	if val != "1:2:3" {
		t.Errorf("c, expected:1:2:3 actual:%v", val)
	}
}

func TestLtsvParseInvalid(t *testing.T) {
	parser, err := NewLtsvParser()
	if err != nil {
		t.Error(err)
	}

	data, err := parser.Parse("a:foo\tb")

	val := data.Data.(map[string]string)["a"]
	if val != "foo" {
		t.Errorf("a, expected:foo actual:%v", val)
	}
	val, ok := data.Data.(map[string]string)["b"]
	if !ok || val != "" {
		t.Errorf("b, expected: actual:%v", val)
	}
}
