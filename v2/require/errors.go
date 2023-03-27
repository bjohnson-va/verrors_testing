package verrors_require

import (
	"testing"

	verrors_assert "github.com/bjohnson-va/verrors_testing/v2/assert"
	"github.com/vendasta/gosdks/verrors"
)

func ErrorTypeEqual(t *testing.T, expectedType verrors.ErrorType, actualErr error) {
	t.Helper()
	if !verrors_assert.ErrorTypeEqual(t, expectedType, actualErr) {
		t.FailNow()
	}
}

func ErrorTypesMatch(t *testing.T, expectedErr error, actualErr error) {
	t.Helper()
	if !verrors_assert.ErrorTypesMatch(t, expectedErr, actualErr) {
		t.FailNow()
	}
}

func NoError(t *testing.T, actualErr error) {
	t.Helper()
	if !verrors_assert.NoError(t, actualErr) {
		t.FailNow()
	}
}
