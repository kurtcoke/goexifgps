package GoExifGPS

import "testing"
    //So RefFormat is in Format.go
func TestRefFormat( t *testing.T) {
	// RefFormat takes (ref string, decGPS float32)
	// if ref is "N" or "E" then decGPS stays the same
	// but if ref is Anything else then decGPS becomes negative.
	// decGPS becomes negative ( decGPS is multiplied by -1 of float32 type)
    DecGPS := float32(1.0)  
    
	Nref := "N"
	//Noref = ''  // try this with nil as decGPS float32 value.
	//Wref = "W" // We know what this will do
	
	RefFormResult := RefFormat(Nref, DecGPS)

	if RefFormResult != DecGPS {
		t.Error("The result should be unchanged, should stay positive.")
	}
}


