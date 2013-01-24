package goexifgps

// Author: kurtcc on github
import (
	"fmt"
	"github.com/rwcarlsen/goexif/tiff"
)
// Get values written as big.Rat returns the result in a more sensible formatting.
func FormatGPS(t *tiff.Tag) float32 {

	Dec := make([]float32, 3)
	for count := 0; count < 3; count++ {

		Numer, Denom := t.Rat2(count) 
		Dec[count] = float32(Numer) / float32(Denom)
	}

	DecGPS := float32(Dec[0]) + float32(Dec[1]/60) + float32(Dec[2]/3600)
	return DecGPS
	//DecGPS := Hours + Minutes/float32(60) + Seconds/float32(3600)
}
//Pass it Lat or Long Ref values and result of formatGPS    
// For Latitude: If N the result is positive if S then negative.
// For Longitude: If the result is E then it is positive, if it is W then it is negative.
func RefFormat(ref string, decGPS float32) float32 { 
	switch ref {
	case "N":
		break
	case "E":
		break

	case "S", "W":
		decGPS *= float32(-1)

	}
	return decGPS
}
 	//Takes values of latitude and longitude returned by RefFormat
	// and turns into a string that is in decimal and can be used on google maps
	// and thus is "map friendly"     
                      
func MapFriendly(latt, lonn float32) string {

	return fmt.Sprintf("%v,%v", latt, lonn)
}
