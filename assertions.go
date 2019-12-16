package verrors_testing

import (
	"github.com/vendasta/gosdks/verrors"
	"testing"
)

func AssertErrorTypeMatches(t *testing.T, expectedType verrors.ErrorType, actualErr error) {
	if actualErr == nil {
		t.Errorf("Expected %s error but got nil", expectedType.String())
		return
	}
	verr := verrors.FromError(actualErr)
	if verr.ErrorType() != expectedType {
		t.Errorf("Expected %s but got %s (%s)", expectedType.String(), verr.ErrorType().String(), verr.Error())
		return
	}
}

func AssertErrorTypesMatch(t *testing.T, expectedErr error, actualErr error) {
	if actualErr == nil && expectedErr != nil {
		t.Errorf("Expected %s error but got nil", expectedErr.Error())
		return
	}
	if actualErr != nil && expectedErr == nil {
		t.Errorf("Expected nil error but got %s", actualErr.Error())
		return
	}
	verr := verrors.FromError(actualErr)
	ex := verrors.FromError(expectedErr)
	if verr.ErrorType() != ex.ErrorType() {
		t.Errorf("Expected %s but got %s (%s)", ex.ErrorType().String(), verr.ErrorType().String(), verr.Error())
		return
	}
}

func AssertVErrorIsNil(t *testing.T, actualErr error) {
	if actualErr == nil {
		return
	}
	verr := verrors.FromError(actualErr)
	t.Errorf("Expected nil but got %s (%s)", verr.ErrorType().String(), verr.Error())
	return
}
