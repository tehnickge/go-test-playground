package steps

import (
	"testing"
)

func TestEmailValidation(t *testing.T) {
	table := []struct {
		email    string
		expected bool
	}{
		{"test@mail.ru", true},
		{"invalid-email", false},
	}

	for _, testCase := range table {
		res := emailValidation(testCase.email)
		if testCase.expected && !res {
			t.Errorf("Expected email %s to be valid, but got invalid", testCase.email)
		}
	}

}
