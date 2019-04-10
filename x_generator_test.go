package xGenerator

import (
	"testing"
	"unicode/utf8"
)

func Test_xGenerator(t *testing.T) {
	length := 10
	xGen := NewXGenerator()
	xID := xGen.Get()

	if utf8.RuneCountInString(xID) != length {
		t.Errorf("%s should consis of %d characters", xID, length)
	}
}
