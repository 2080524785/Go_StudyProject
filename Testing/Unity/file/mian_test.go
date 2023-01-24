package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
func TestProcessFirstLine(t *testing.T) {
	firstLine := ProcessFirstLine()
	println(firstLine)
	assert.Equal(t, "line00", firstLine)
}
