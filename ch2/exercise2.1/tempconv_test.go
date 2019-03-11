package tempconv

import (
	"testing"
)

func TestName(t *testing.T) {
	var c Celsius = 100
	t.Log(CToK(c))
	t.Log(CToF(c))

	var f Fahrenheit = 212
	t.Log(FToK(f))
	t.Log(FToC(f))
}
