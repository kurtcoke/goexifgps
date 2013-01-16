package goexifgps

import "testing"

func TestGetGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	// Check that TrueImage returns no error

	TrueI, err := OpenClose(TrueImage)
	// error checking
	if err != nil {
		t.Error("This image should open and decode fine.")
	}
	
	_, err = GetGPS(TrueI)
	// If err != nil something is wrong
	if err != nil {
		t.Error("This image should work without returning errors!")
	}
	
}
