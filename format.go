package goexifgps

// Author: kurtcc on github
import (
	"fmt"
	"github.com/rwcarlsen/goexif/tiff"
)

//This will be called format.go
// This part will only be used by tage GPSlatitude and GPSLongitude
//Usage example 
// 
func FormatGPS(t *tiff.Tag) float32 {
	//I could also make this float64 but I don't think this would give better results?   
	Dec := make([]float32, 3)
	for count := 0; count < 3; count++ {
		/* I could just not use all this bload if i used Rat2 method from tiff package lets test if first. 	
		A := t.Rat(count)
		Numer := A.Num()
		Denom := A.Denom()
		Dec[count] = float32(Numer) / float32(Denom) */
		Numer,Denom := t.Rat2(count) // They are now of int64 type
		Dec[count] = float32(Numer)/float32(Denom) //Still need to type cast
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
	// So only two left will always only be one of them will use this method twice
	// once for Latitude once for Longitude
	default:
		decGPS *= float32(-1)

	}
	return decGPS
}
func MapFriendly(latt, lonn float32) string {
	//Takes values of latitude and longitude returned by RefFormat
	// and turns into a string that is in decimal and can be used on google maps
	// and thus is "map friendly"     
	//Think about rewriting everything with type float64 for better precision
	// if it will help   

	return fmt.Sprintf("%v,%v", latt, lonn)
}
