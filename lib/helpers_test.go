package lib

import (
	"testing"
)

func TestStringPtr(t *testing.T) {
	str := "test"
	strPtr := StringPtr(str)
	if *strPtr != str {
		t.Errorf("Expected %s, got %s", str, *strPtr)
	}
}

func TestStringToInt(t *testing.T) {
	table := []struct {
		input    string
		expected int
		valid    bool
	}{
		{"123", 123, true},
		{"-45", -45, true},
		{"0", 0, true},
		{"test", 0, false},
	}
	for _, testCase := range table {
		result, err := StringToInt(testCase.input)
		if testCase.valid {
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", testCase.input, err)
			}
			if result != testCase.expected {
				t.Errorf("Expected %d for input %s, got %d", testCase.expected, testCase.input, result)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error for input %s, got nil", testCase.input)
			}
		}
	}

	for _, tc := range table {
		result, err := StringToInt(tc.input)
		if tc.valid {
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tc.input, err)
			}
			if result != tc.expected {
				t.Errorf("Expected %d for input %s, got %d", tc.expected, tc.input, result)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error for input %s, got nil", tc.input)
			}
		}
	}
}
