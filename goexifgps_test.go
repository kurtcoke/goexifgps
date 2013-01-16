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
	FalseImage := "testdata/WTgX4.jpg"
	ExifData2, err := OpenClose(FalseImage)
	//This should work fine still, but below should return error.
	_, err = DecodeGPS(ExifData2)
	if err == nil {
		t.Error("This image has been tested and should return an error!", err)
	}
}
