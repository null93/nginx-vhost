package vhost

import (
	"testing"
)

func TestParsing(t *testing.T) {
	_, errValid := ParseInputSchema(
		[]byte(`
            foo:
                pattern: yes-no
                description: foo
                value: foo
                array_separator: ","
                provisioner_only: true
            bar:
                custom_pattern: ^bar$
        `),
	)
	if errValid != nil {
		t.Fatal(errValid)
	}
	_, errBlank := ParseInputSchema([]byte(``))
	if errBlank != nil {
		t.Fatal(errBlank)
	}
	_, errReserved := ParseInputSchema([]byte(`site_name: { pattern: "yes-no" }`))
	if errReserved != ErrReservedKey {
		t.Fatal(errReserved)
	}
}

func TestValidation(t *testing.T) {
	errMissingPattern := ValidateDefinition(Definition{})
	if errMissingPattern != ErrMissingPattern {
		t.Fatal(errMissingPattern)
	}
	errMultiplePatterns := ValidateDefinition(Definition{
		Pattern:       "yes-no",
		CustomPattern: "^bar$",
	})
	if errMultiplePatterns != ErrMultiplePatterns {
		t.Fatal(errMultiplePatterns)
	}
	errValid := ValidateDefinition(Definition{
		Pattern: "yes-no",
	})
	if errValid != nil {
		t.Fatal(errValid)
	}
	errValid = ValidateDefinition(Definition{
		CustomPattern: "^bar$",
	})
	if errValid != nil {
		t.Fatal(errValid)
	}
	errInvalidPattern := ValidateDefinition(Definition{
		Pattern: "baz",
	})
	if errInvalidPattern != ErrInvalidPattern {
		t.Fatal(errInvalidPattern)
	}
	errValidCustomPattern := ValidateDefinition(Definition{
		CustomPattern: "^bar$",
	})
	if errValidCustomPattern != nil {
		t.Fatal(errValidCustomPattern)
	}
	errInvalidCustomPattern := ValidateDefinition(Definition{
		CustomPattern: "bar(",
	})
	if errInvalidCustomPattern != ErrInvalidCustomPattern {
		t.Fatal(errInvalidCustomPattern)
	}
}
