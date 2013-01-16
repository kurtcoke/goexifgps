package goexifgps

import "testing"

func TestGetGPS(t *testing.T) {
	TrueImage := "testdata/_JEF018993_sm.jpg"
	FalseImage := "testdata/WTgX4.jpg"
	// Check that TrueImage returns no error
	// Check that FalseImage returns an error

	TrueI, err := OpenClose(TrueImage)
	// error checking
	if err != nil {
		t.Error("This image should open and decode fine.")
	}
	FalseI, err := OpenClose(FalseImage)
	if err != nil {
		t.Error("This image should open and decode fine.")
	}
	_, err := GetGPS(TrueI)
	// If err != nil something is wrong
	if err != nil {
		t.Error("This image should work without returning errors!")
	}
	_, err := GetGPS(FalseI)
	// If err == nil something is wrong.
	if err == nil {
		t.Error("This image should return an error!")
	}
}
