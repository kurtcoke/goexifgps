package GoExifGPS

import "testing"

func TestContainsGPS( t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	if ContainsGPS(TrueImage) != true {
		t.Error("Image has GPS data, this should work fine!")
	}
  FalseImage := "testdata/WTgX4.jpg"
  if ContainsGPS(FalseImage) == true {
	  t.Error("Image does not contain gps data!")
  }
}
