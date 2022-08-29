package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

var netropy = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewULID() string {
	ms := ulid.Timestamp(time.Now())

	return fmt.Sprintf("%s", ulid.MustNew(ms, netropy))
}
