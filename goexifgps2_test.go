package goexifgps

import "testing"

func TestDecodeGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	//Image does contain GPS data.
	//FalseImage := "testdata/WTgX4.jpg"
	// Does not contain GPS data.
	// First will return string, latter will return error.
	ExifData, err := OpenClose(TrueImage)
	// We know the above works, we have GoExifGPS2_test.go to prove it.
	if err != nil {
		t.Error("This should work without returning an error.")
	}
}
