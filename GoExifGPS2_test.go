package goexifgps

import "testing"

func TestOpenClose(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	// This image is working and should open with no errors.
	_, err := OpenClose(TrueImage)
	if err != nil {
		t.Error("This image has been tested and should work.", err)
	}
}
