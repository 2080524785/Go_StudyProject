package example

import (
	"github.com/awnumar/fastrand"
)

func FastSelect() int {
	return ServerIndex[fastrand.Intn(10)]
}
