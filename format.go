package goexifgps

// Author: kurtcc on github
import (
	"fmt"
	"github.com/rwcarlsen/goexif/tiff"
)

func FormatGPS(t *tiff.Tag) float32 {

	Dec := make([]float32, 3)
	for count := 0; count < 3; count++ {

		Numer, Denom := t.Rat2(count) // They are now of int64 type
		Dec[count] = float32(Numer) / float32(Denom)
	}

	DecGPS := float32(Dec[0]) + float32(Dec[1]/60) + float32(Dec[2]/3600)
	return DecGPS
	//dGPS := Hours + Minutes/float32(60) + Seconds/float32(3600)
}

func RefFormat(ref string, decGPS float32) float32 { //Pass Ref and result of formatGPS
	// The program would throw errors if there was no gps data before it ran this part!
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
func MapFriendly(latt, lonn float32) string {
	//Takes values of latitude and longitude returned by RefFormat
	// and turns into a string that is in decimal and can be used on google maps
	// and thus is "map friendly"     

	return fmt.Sprintf("%v,%v", latt, lonn)
}
