package goexifgps

// Author : kurtcc on github
// This will be called parse.go
import (
	"encoding/json"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
	"os"
	"strings"

//This is how to print a given field
)

type Exif struct {
	tif *tiff.Tiff

	main map[FieldName]*tiff.Tag
}
type FieldName string

type GeoFields struct {
	LatRef, Lat, LongRef, Long string
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func TrimPrefix(s string) string {
	s = s[1:]
	return s
}

// Use like this
//LatRef, Lat, LongRef,Longd := OpenParseJson("_JEF018993_sm.jpg") 
func OpenClose(filename string) (*exif.Exif, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	ExifData, err2 := exif.Decode(f)
	f.Close()
	if err2 != nil {
		return nil, err2
	}
	return ExifData, nil

}

// Gonna make it also return errors. (*GeoFields, error)
func GetGPS(E *exif.Exif) (*GeoFields, error) {
	// I want this to return all four values each as a string.	
	// Was named OpenParseJson now named GetGPS
	// Gebruik exif.Get[Some field related to gps] om te check vir errors
	F := new(GeoFields)
	LatVal, err := E.Get("GPSLatitude")
	if err != nil {
		panic(err)
	}
	LongVal, err := E.Get("GPSLongitude")
	if err != nil {
		panic(err)
	}

	LatRefVal, err := E.Get("GPSLatitudeRef") //Lat and LatRef
	LongRefVal, err := E.Get("GPSLongitudeRef")

	F.LatRef = LatRefVal.StringVal()
	F.LongRef = LongRefVal.StringVal()
	
	Tnum1 := TrimPrefix(LatVal2)
	F.Lat = TrimSuffix(Tnum1, "]")

	// *** Longitude
	Tnum2 := TrimPrefix(LongVal2)
	F.Long = TrimSuffix(Tnum2, "]")

	return F, nil
}
