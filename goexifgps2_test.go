package goexifgps

import "testing"

func TestDecodeGPS(t *testing.T) {
	//TrueImage := "testdata/_JEF018993_sm.jpg"
	//Image does contain GPS data.
	FalseImage := "testdata/WTgX4.jpg"
	// Does not contain GPS data.
	// Should return an error since exif.Get("GPSLatitude") will return an error if no
	// data for given field is available.
	 _ , err:= OpenClose(FalseImage)
	if err == nil {
		t.Error("This image should return an error!")
		
	}
}
