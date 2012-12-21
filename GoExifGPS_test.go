package GoExifGPS

import "testing"

func TestContainsGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	YesGPS, _ := ContainsGPS(TrueImage)
	if YesGPS != true {
		t.Error("Image has GPS data, this should work fine!")
	}
	FalseImage := "testdata/WTgX4.jpg"
	NoGPS, _ := ContainsGPS(FalseImage)
	if NoGPS == true {
		t.Error("Image does not contain gps data!")
	}
}
