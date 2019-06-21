package parse

import (
	"testing"

	"github.com/chaseadamsio/goorgeous/lex"
)

func TestIsLink(t *testing.T) {
	testCases := []struct {
		value    string
		expected bool
	}{
		{"[[https://https://orgmode.org/manual/Link-format.html#Link-format][Org Mode: link format]]", true},
		{"[[https://https://orgmode.org/manual/Link-format.html#Link-format]]", true},
		{"[[https://https://orgmode.org/manual/Link-format.html#Link-format[]Org Mode: link format]]", false},
		{"[[https://https://orgmode.org/manual/Link-format.html#Link-format][Org Mode: link format]", false},
		{"[https://https://orgmode.org/manual/Link-format.html#Link-format][Org Mode: link format]]", false},
		{"[https://https://orgmode.org/manual/Link-format.html#Link-format]]", false},
	}

	for _, tc := range testCases {
		var items []lex.Item
		lexedItems := lex.NewLexer(tc.value)
		for item := range lexedItems {
			items = append(items, item)
		}
		p := parser{
			items: items,
		}
		if found, _ := p.matchesLink(0); found != tc.expected {
			t.Errorf("expected \"%s\" to be %t", tc.value, tc.expected)
		}
	}
}
