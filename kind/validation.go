package kind

import (
	"fmt"
)

func stringIsValidToml(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
		return warnings, errors
	}
	_, err := normalizeToml(v)
	if err != nil {
		errors = append(errors, fmt.Errorf("%s is not valid toml: %s", k, err))
	}
	return warnings, errors
}
