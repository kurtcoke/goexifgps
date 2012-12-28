package GoExifGPS

import "testing"

func TestRefFormat( t *testing.T) {
	// RefFormat takes (ref string, decGPS float32)
	// if ref is "N" or "E" then decGPS stays the same
	// but if ref is Anything else then decGPS becomes negative.
	// decGPS becomes negative ( decGPS is multiplied by -1 of float32 type)

