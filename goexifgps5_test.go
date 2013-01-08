package goexifgps

import "testing"
import "fmt"

func TestDecodeGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	//Image does contain GPS data.
	//FalseImage := "testdata/WTgX4.jpg"
	// Does not contain GPS data.
	// First will return string, latter will return error.
	ExifData, _ := OpenClose(TrueImage)
	// We know the above works, we have GoExifGPS2_test.go to prove it.
	F := GetGPS(ExifData)
	fmt.Println(F.Lat)

}
