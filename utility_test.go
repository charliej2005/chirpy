package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello mother kerfuffle",
			expected: "hello mother ****",
		},
		{
			input:    "nice to meet you",
			expected: "nice to meet you",
		},
		{
			input:    "sharbert for brains",
			expected: "**** for brains",
		},
		{
			input:    "shut the fornax up",
			expected: "shut the **** up",
		},
	}

	blocked := make(map[string]struct{})
	blocked["kerfuffle"] = struct{}{}
	blocked["sharbert"] = struct{}{}
	blocked["fornax"] = struct{}{}

	for _, c := range cases {
		actual := cleanInput(c.input, blocked)
		if len(c.expected) != len(actual) {
			t.Errorf("lengths do not match | expected: %v, actual: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("words do not match | index: %v, expected: %v, actual: %v", i, expectedWord, word)
			}
		}
	}
}
