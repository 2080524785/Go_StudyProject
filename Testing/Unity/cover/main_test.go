package cover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgePassLineTrue(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}
func TestJudgePassLineFalse(t *testing.T) {
	isPass := JudgePassLine(50)
	assert.Equal(t, false, isPass)
}
