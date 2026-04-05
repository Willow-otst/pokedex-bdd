
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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "One",
			expected: []string{"one"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %v, expected %v",
				 c.input, actual, c.expected)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) returned %v, expected %v",
					 c.input, actual, c.expected)
				break
			}
		}
	}
}
