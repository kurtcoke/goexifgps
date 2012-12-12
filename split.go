package format

import (
	"fmt"
	"strconv"
	"strings"
)

//This will be called format.go

func formatGPS(s string) float32 {
	// Accepts string and returns float32 representation.
	// Then we can output this to file or stdout and use it on google maps (Correct format.)

	Fs := strings.Split(s, " ") // Formatted s
	// Fs was foo previously

	sHours := strings.Split(Fs[0], "/")   //Hours  was splitN
	sMinutes := strings.Split(Fs[1], "/") // Minutes   was splitN2
	sSeconds := strings.Split(Fs[2], "/") //Seconds   was splitN3

	//Hours 
	sp, _ := strconv.Atoi(sHours[0])
	sp2, _ := strconv.Atoi(sHours[1])
	Hours := float32(sp) / float32(sp2)

	//Minutes
	Msp, _ := strconv.Atoi(sMinutes[0])
	Msp2, _ := strconv.Atoi(sMinutes[1])
	Minutes := float32(Msp) / float32(Msp2)

	//Seconds
	Ssp, _ := strconv.Atoi(sSeconds[0])
	Ssp2, _ := strconv.Atoi(sSeconds[1])
	Seconds := float32(Ssp) / float32(Ssp2)

	//Try stackoverflow suggestion D = H + M/60 + 60/3600
	//Lat N==- S==+ 
	//Long E==+ W==-
	dGPS := Hours + Minutes/float32(60) + Seconds/float32(3600)
	//fmt.Println(dGPS)
	//fmt.Println(minutes)
	//fmt.Println(seconds)
	//for xx := range foofoo {
	//fmt.Println(foofoo[xx])
	//}
	return dGPS //Decimal GPS 
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
