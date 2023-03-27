package verrors_assert

import (
	"testing"

	"github.com/bjohnson-va/verrors_testing"
	"github.com/vendasta/gosdks/verrors"
)

func ErrorTypeEqual(t *testing.T, expectedType verrors.ErrorType, actualErr error) bool {
	t.Helper()
	return verrors_testing.AssertErrorTypeEqual(t, expectedType, actualErr)
}

func ErrorTypesMatch(t *testing.T, expectedErr error, actualErr error) bool {
	t.Helper()
	return verrors_testing.AssertErrorTypesMatch(t, expectedErr, actualErr)
}

func NoError(t *testing.T, actualErr error) bool {
	t.Helper()
	return verrors_testing.AssertVErrorIsNil(t, actualErr)
}
