package domain

import "testing"

func Test_newRuleValid(t *testing.T) {
	validRegularExpresion := "pdf$"
	dest := "dest"

	r := NewRule(validRegularExpresion, dest)

	if r.DestinationFolder != dest {
		t.Errorf("Invalid rule, dest expected %s, got %s", dest, r.DestinationFolder)
	}
}

func Test_newRuleInvalid(t *testing.T) {
	invalidRegexp := "["

	//si una regla recibe
	defer func() {

		if r := recover(); r == nil {
			t.Errorf("Panic expected with invalid regexp")
		}
	}()

	NewRule(invalidRegexp, "")
}
