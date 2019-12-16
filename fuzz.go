package verrors_testing

import (
	"github.com/vendasta/gosdks/verrors"
	"math/rand"
	"time"
)

func RandomVError() error {
	// TODO: More futureproof implementation (i.e. handle new additions)
	rand.Seed(time.Now().Unix())
	t := rand.Int63n(int64(verrors.BadGateway))
	return verrors.New(verrors.ErrorType(t), "Random vError generated by verrors_testing")
}
