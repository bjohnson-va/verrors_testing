package verrors_testing

import (
	"math/rand"
	"testing"
	"time"

	"github.com/vendasta/gosdks/verrors"
)

func RandomVError() verrors.ServiceError {
	// TODO: More futureproof implementation (i.e. handle new additions)
	rand.Seed(time.Now().Unix())
	t := rand.Int63n(int64(verrors.BadGateway))
	return verrors.New(verrors.ErrorType(t), "Random vError generated by verrors_testing")
}

func RandomVErrorExcept(not ...verrors.ErrorType) verrors.ServiceError {
Choose:
	for {
		// TODO: There's almost certainly a way to do this that doesn't involve an infinite loop of retries
		err := RandomVError()
		for _, e := range not {
			if err.ErrorType() == e {
				continue Choose
			}
		}
		return err
	}
}

func FuzzRandomVError(
	f *testing.F,
	testFn func(
		t *testing.T,
		vErrorCase verrors.ServiceError,
	),
) {
	errors := []verrors.ErrorType{
		verrors.NotFound,
		verrors.InvalidArgument,
		verrors.AlreadyExists,
		verrors.PermissionDenied,
		verrors.Unauthenticated,
		verrors.Unimplemented,
		verrors.Unknown,
		verrors.Internal,
		verrors.Gone,
		verrors.Unavailable,
		verrors.FailedPrecondition,
		verrors.DeadlineExceeded,
		verrors.ResourceExhausted,
		verrors.Aborted,
		verrors.Canceled,
		verrors.BadGateway,
	}
	for _, vr := range errors {
		f.Add(int(vr))
	}
	f.Fuzz(func(
		t *testing.T,
		errCode int,
	) {
		err := verrors.New(verrors.ErrorType(errCode), "error defined in fuzz test setup [verrors_testing]")
		testFn(t, err)
	})
}

func FuzzRandomVErrorExcept(
	f *testing.F,
	testFn func(
		t *testing.T,
		vErrorCase verrors.ServiceError,
	),
	except verrors.ErrorType,
) {
	errors := []verrors.ErrorType{
		verrors.NotFound,
		verrors.InvalidArgument,
		verrors.AlreadyExists,
		verrors.PermissionDenied,
		verrors.Unauthenticated,
		verrors.Unimplemented,
		verrors.Unknown,
		verrors.Internal,
		verrors.Gone,
		verrors.Unavailable,
		verrors.FailedPrecondition,
		verrors.DeadlineExceeded,
		verrors.ResourceExhausted,
		verrors.Aborted,
		verrors.Canceled,
		verrors.BadGateway,
	}
	for _, vr := range errors {
		if vr == except {
			continue
		}
		f.Add(int(vr))
	}
	f.Fuzz(func(
		t *testing.T,
		errCode int,
	) {
		err := verrors.New(verrors.ErrorType(errCode), "error defined in fuzz test setup [verrors_testing]")
		testFn(t, err)
	})
}
