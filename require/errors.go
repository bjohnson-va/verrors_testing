package verrors_require

import (
	"testing"

	"github.com/bjohnson-va/verrors_testing"
	verrors_assert "github.com/bjohnson-va/verrors_testing/assert"
	"github.com/vendasta/gosdks/verrors"
)

func ErrorTypeEqual(t *testing.T, expectedType verrors.ErrorType, actualErr error) {
	if !verrors_assert.ErrorTypeEqual(t, expectedType, actualErr) {
		t.FailNow()
	}
}

func ErrorTypesMatch(t *testing.T, expectedErr error, actualErr error) {
	if !verrors_testing.AssertErrorTypesMatch(t, expectedErr, actualErr) {
		t.FailNow()
	}
}

func NoError(t *testing.T, actualErr error) {
	if !verrors_testing.AssertVErrorIsNil(t, actualErr) {
		t.FailNow()
	}
}
