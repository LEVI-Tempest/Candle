package utils

import "fmt"

// Aggregate aggregates multiple errors into one error
func Aggregate(errs []error) error {
	var errStr string
	var nonNilErrs []error

	// filter out nil errors
	for _, err := range errs {
		if err != nil {
			nonNilErrs = append(nonNilErrs, err)
		}
	}

	// if no errors, return nil
	if len(nonNilErrs) == 0 {
		return nil
	}

	// aggregate errors
	errStr = fmt.Sprintf("%d errs.\n", len(nonNilErrs))
	for i, err := range nonNilErrs {
		errStr += fmt.Sprintf("\terr[%d]: %v\n", i, err)
	}
	return fmt.Errorf(errStr)
}
