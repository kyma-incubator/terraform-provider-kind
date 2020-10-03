package kind

import (
	"strings"
	"testing"
)

func TestStringIsValidToml(t *testing.T) {
	cases := []struct {
		Name             string
		Value            interface{}
		Key              string
		ExpectedErrors   int
		ExpectedWarnings int
	}{
		{
			Name:           "PassingNonStringIsAnError",
			Value:          struct{}{},
			ExpectedErrors: 1,
		},
		{
			Name:  "ValidTomlIsValid",
			Value: "the_answer_to_everything = 42",
		},
		{
			Name:           "NilIsInvalid",
			Value:          nil,
			ExpectedErrors: 1,
		},
		{
			Name:  "EmptyStringIsValid",
			Value: "",
		},
		{
			Name:           "InvalidTomlIsInvalid",
			Value:          strings.Join([]string{"fruits = []", "[[fruits]]"}, "\n"),
			ExpectedErrors: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			warnings, errors := stringIsValidToml(tc.Value, tc.Key)
			if len(warnings) != tc.ExpectedWarnings {
				t.Errorf("expected %d warnings but got len(%v) = %d", tc.ExpectedWarnings, warnings, len(warnings))
			}
			if len(errors) != tc.ExpectedErrors {
				t.Errorf("expected %d warnings but got len(%v) = %d", tc.ExpectedErrors, errors, len(errors))
			}
		})
	}
}
