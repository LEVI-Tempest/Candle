package utils

import (
	"fmt"
	"testing"
)

func TestAggregate(t *testing.T) {
	errs := []error{
		fmt.Errorf("err1"),
		fmt.Errorf("err2"),
		fmt.Errorf("err3"),
	}
	err := Aggregate(errs)
	fmt.Println(err)
}
