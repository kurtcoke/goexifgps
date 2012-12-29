package GoExifGPS

import "testing"

func TestDecodeGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	//Image does contain GPS data.
	//FalseImage := "testdata/WTgX4.jpg"
	// Does not contain GPS data.
	// First will return string, latter will return error.
	ExifData, _ := OpenClose(TrueImage)
	// We know the above works, we have GoExifGPS2_test.go to prove it.
	GPSLocation := DecodeGPS(ExifData)
	if len(GPSLocation) == 0 {
		t.Error("This image should open fine with no errors!")
		// Should return a string value that can be used to find location on 
		// google maps
	}
}
