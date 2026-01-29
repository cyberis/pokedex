package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "basic test",
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "empty input",
			input:    "     ",
			expected: []string{},
		},
		{
			name:     "mixed case and spaces",
			input:    " GoLang  is   FUN ",
			expected: []string{"golang", "is", "fun"},
		},
		{
			name:     "single word",
			input:    "   Test   ",
			expected: []string{"test"},
		},
		{
			name:     "multiple spaces between words",
			input:    "This    is  a   test",
			expected: []string{"this", "is", "a", "test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanInput(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("cleanInput(%q) = %#v; want %#v", tt.input, result, tt.expected)
			}
		})
	}
}
