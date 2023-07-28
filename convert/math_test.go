package convert_test

import (
	"github.com/westfruit/kcc-toolkit/convert"
	"testing"
)

func TestXxx(t *testing.T) {
	f := 0.336
	f = f * 100
	v := convert.KeepDecimal(f, 0)

	t.Log(v)
}
