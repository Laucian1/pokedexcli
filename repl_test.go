package main

import "testing"

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
			input:    "GOTTA catch 'Em ALL!",
			expected: []string{"gotta", "catch", "'em", "all!"},
		},
		{
			input:    "I want to be the very best",
			expected: []string{"i", "want", "to", "be", "the", "very", "best"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %s, but got %s", expectedWord, word)
			}
		}
	}
}
