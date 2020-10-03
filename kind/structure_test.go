package kind

import (
	"testing"
)

func TestNormalizeTomlString(t *testing.T) {
	cases := []struct {
		Name            string
		Input           string
		ExpectedOutput  string
		ExpectUnchanged bool
		ExpectError     bool
	}{
		{
			Name: "WellFormattedIsReturnedAsIs",
			Input: `name = "Test"

[section]
  key = "value"
`,
			ExpectUnchanged: true,
		},
		{
			Name: "MalformedInputResultsInErrorAndReturnsInput",
			Input: `fruit = []
[[fruit]]
`,
			ExpectUnchanged: true,
			ExpectError:     true,
		},
		{
			Name: "UnformattedInputIsFormatted",
			Input: `name = "test"
[fruit.apple]
[animal]
[fruit]
`,
			ExpectedOutput: `name = "test"

[animal]

[fruit]

  [fruit.apple]
`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			output, err := normalizeToml(tc.Input)
			if err != nil {
				if !tc.ExpectError {
					t.Error("received error but expected no errors for case", err)
				}
			}
			if tc.ExpectUnchanged && output != tc.Input {
				t.Errorf("received:\n---\n%s\n---\n but expected input \n---\n%s\n---\n to be unchanged", output, tc.Input)
			}
			if !tc.ExpectUnchanged && output != tc.ExpectedOutput {
				t.Errorf("received \n---\n%s\n---\n but expected \n---\n%s\n---\n", output, tc.ExpectedOutput)
			}
		})
	}
}
