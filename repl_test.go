package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "     	",
			expected: []string{},
		},
		{
			input: "HEY thEre",
			expected: []string{"hey", "there"},
		},
		{
			input: "!)#*",
			expected: []string{"!)#*"},
		},
		{
			input: "   Let's TEST         this, SHall wE?     ",
			expected: []string{"let's", "test", "this,", "shall", "we?"},
		},
	}

	for _, testCase := range cases{
		output := cleanInput(testCase.input)
		if len(output) != len(testCase.expected){
			t.Errorf("Length of output: %v, expected: %v", len(output), len(testCase.expected))
			continue
		}
		for i, word := range output{
			if testCase.expected[i] != word{
				t.Errorf("Output is not as expected: %s != %s", word, testCase.expected[i])
				break
			}
		}
	}
}