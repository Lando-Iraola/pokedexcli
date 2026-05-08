package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  charizar     spirigatito",
			expected: []string{"charizar", "spirigatito"},
		},
		{
			input:    "  charizar     spirigatito                                      dodrio",
			expected: []string{"charizar", "spirigatito", "dodrio"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}
	for _, c := range cases {

		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Number of expected results: %d, results got: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word %s doesn't match expected: %s", word, expectedWord)
			}
		}
	}
}
