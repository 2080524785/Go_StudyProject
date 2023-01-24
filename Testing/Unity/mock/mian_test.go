package file

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
func TestProcessFirstLine(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	firstLine := ProcessFirstLine()
	assert.Equal(t, "line000", firstLine)
}
