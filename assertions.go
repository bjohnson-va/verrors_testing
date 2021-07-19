// Deprecated: Use verrors_assert
package verrors_testing

import (
	"fmt"
	"testing"

	"github.com/vendasta/gosdks/verrors"
)

func AssertErrorTypeEqual(t *testing.T, expectedType verrors.ErrorType, actualErr error) bool {
	if actualErr == nil {
		t.Errorf("Expected %s error but got nil", expectedType.String())
		return false
	}
	verr := verrors.FromError(actualErr)
	if verr.ErrorType() != expectedType {
		t.Errorf("Expected %s but got %s (%s)", expectedType.String(), verr.ErrorType().String(), verr.Error())
		return false
	}
	return true
}

func AssertErrorTypesMatch(t *testing.T, expectedErr error, actualErr error) bool {
	if actualErr == nil && expectedErr == nil {
		return false
	}
	if actualErr == nil && expectedErr != nil {
		t.Errorf("Expected %s error but got nil", expectedErr.Error())
		return false
	}
	if actualErr != nil && expectedErr == nil {
		t.Errorf("Expected nil error but got %s", actualErr.Error())
		return false
	}
	verr := verrors.FromError(actualErr)
	ex := verrors.FromError(expectedErr)
	if verr.ErrorType() != ex.ErrorType() {
		t.Errorf(
			"\n%10s: %20s (%s)\n"+
				"%10s: %20s (%s)",
			"Expected", ex.ErrorType().String(), ex.Error(),
			"Actual", verr.ErrorType().String(), verr.Error(),
		)
		return false
	}
	return true
}

func AssertVErrorIsNil(t *testing.T, actualErr error) bool {
	if actualErr == nil {
		return true
	}
	verr := verrors.FromError(actualErr)
	Fail(
		t, fmt.Sprintf(
			`Expected nil but got:\n%s ["%s", "%s"]`,
			verr.ErrorType().String(), verr.Error(), verr.GetInternalMessage(),
		),
	)
	return false
}
