package goexifgps

// Author: kurtcc on github
import (
	"fmt"
	"strconv"
	"strings"
)

//This will be called format.go

func FormatGPS(t *tiff.Tag) float64 {
     
	 Dec := (make[]float64,3)
	 for count:= 0; count <3; count++ {
		 A := t.Rat(count)
		 Numer := A.Num()
		 Denom := A.Denom()
		 Dec[count] := float64(Numer)/float64(Denom)
		 switch count {
			 case 0:
			 	break
			case 1:
			Dec[1] /= 60
			case 2:
			Dec[2] /= 3600
		 }
		 return Dec[count]
	//dGPS := Hours + Minutes/float32(60) + Seconds/float32(3600)
}
return Dec[0] + Dec[1] + Dec[2]
}

//Add function to for Lat and LongRef to make negative if S or W
//LongitudeRef and LatitudeRef has type string.
func RefFormat(ref string, decGPS float32) float32 { //Pass Ref and result of formatGPS

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
