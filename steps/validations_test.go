package steps

import (
	"testing"
)

func TestEmailValidation(t *testing.T) {
	var email string = "test@mail.ru"
	isValid := emailValidation(email)
	if !isValid {
		t.Errorf("Expected email to be valid, got invalid")
	}
}
