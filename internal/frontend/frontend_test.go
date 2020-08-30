package frontend

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Chdir("../..")
	os.Exit(m.Run())
}
