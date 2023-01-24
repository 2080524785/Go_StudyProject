package example

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	if output != expectOutput {
		t.Errorf("Expected %s do not match actual %s", expectOutput, output)
	}

}
