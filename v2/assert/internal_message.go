package verrors_assert

import (
	"strings"
	"testing"

	"github.com/vendasta/gosdks/verrors"
)

func InternalMessageContains(t *testing.T, err error, expectedContent string) bool {
	t.Helper()
	vErr := verrors.FromError(err)
	internalMessage := vErr.GetInternalMessage()
	if !strings.Contains(internalMessage, expectedContent) {
		t.Errorf(
			"\nExpected Internal Message to contain substring"+
				"\nInternal Message: %s"+
				"\nSubString: %s",
			internalMessage, expectedContent,
		)
		return false
	}
	return true
}
