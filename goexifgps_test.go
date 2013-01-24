package goexifgps

import "testing"

func TestOpenClose(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	// This image is working and should open with no errors.
	ExifData, err := OpenClose(TrueImage)
	_, err = DecodeGPS(ExifData)
	if err != nil {
		t.Error("This image has been tested and should work.", err)
	}

}
