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

func TestDecodeGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	//Image does contain GPS data.
	//FalseImage := "testdata/WTgX4.jpg"
	// Does not contain GPS data.
	// First will return string, latter will return error.
	_, err := OpenClose(TrueImage)
	// We know the above works, we have GoExifGPS2_test.go to prove it.
	if err != nil {
		t.Error("This should work without returning an error.")
	}
}

func TestGetGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	// Check that TrueImage returns no errors

	TrueI, err := OpenClose(TrueImage)
	if err != nil {
		t.Error("This image should open and decode fine.")
	}
	_, err = GetGPS(TrueI)
	// If err != nil something is wrong
	if err != nil {
		t.Error("This image should work without returning errors!")
	}
}
