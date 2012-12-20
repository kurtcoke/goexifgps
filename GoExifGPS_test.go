package GoExifGPS

import "testing"

func TestContainsGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	YesGPS, err := ContainsGPS(TrueImage)
	if YesGPS != true {
		t.Error("Image has GPS data, this should work fine!")
	}
	FalseImage := "testdata/WTgX4.jpg"
	NoGPS, err2 := ContainsGPS(FalseImage)
	if NoGPS == true {
		t.Error("Image does not contain gps data!")
	}
}
