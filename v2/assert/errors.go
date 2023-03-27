package verrors_assert

import (
	"fmt"
	"testing"
	"github.com/bjohnson-va/verrors_testing/v2"
	"github.com/vendasta/gosdks/verrors"
)

func ErrorTypeEqual(t *testing.T, expectedType verrors.ErrorType, actualErr error) bool {
	t.Helper()
	if actualErr == nil {
		t.Errorf("Expected %s error but got nil", expectedType.String())
		return false
	}
	verr := verrors.FromError(actualErr)
	if verr.ErrorType() != expectedType {
		t.Errorf(
			"Expected %s but got %s (%s) []",
			expectedType.String(),
			verr.ErrorType().String(),
			verr.Error(),
			verr.GetInternalMessage(),
		)
		return false
	}
	return true
}

func ErrorTypesMatch(t *testing.T, expectedErr error, actualErr error) bool {
	t.Helper()
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
				"%10s: %20s (%s) [%s]",
			"Expected", ex.ErrorType().String(), ex.Error(),
			"Actual", verr.ErrorType().String(), verr.Error(), verr.GetInternalMessage(),
		)
		return false
	}
	return true
}

func NoError(t *testing.T, actualErr error) bool {
	t.Helper()
	if actualErr == nil {
		return true
	}
	verr := verrors.FromError(actualErr)
	verrors_testing.Fail(
		t, fmt.Sprintf(
			`Expected nil but got:
{type: %s, msg: "%s", internalMsg: "%s"}`,
			verr.ErrorType().String(),
			verr.Error(),
			verr.GetInternalMessage(),
		),
	)
	return false
}
