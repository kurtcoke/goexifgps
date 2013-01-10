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
	var ExifData *exif.Exif
	f, err := os.Open(filename)
	defer f.Close()
	if err == nil {
		ExifData, err = exif.Decode(f)
 
	}
	return nil, err
}

// Gonna make it also return errors. (*GeoFields, error)
func GetGPS(E *exif.Exif) (*GeoFields, error) {
	// I want this to return all four values each as a string.	
	// Was named OpenParseJson now named GetGPS
	// Gebruik exif.Get[Some field related to gps] om te check vir errors
	F := new(GeoFields)
	_, err := E.Get("GPSLatitude")
	if err != nil {
		return F, err
	}

	b, err2 := E.MarshalJSON()
	if err2 != nil {
		panic(err2) //Format die output properly
	}
	var dat map[string]interface{}
	if err = json.Unmarshal(b, &dat); err != nil {
		panic(err)
	}
    // This is the last commit where I use dat["GPSLatitude"] etc, everything works fine here
	// next I need to try just use x.Get("GPSLatitude")
	// if I can't get that to work I will fall back to this commit where everything worked.
	F.LatRef = dat["GPSLatitudeRef"].(string) //Lat and LatRef
	num := dat["GPSLatitude"].([]interface{})

	F.LongRef = dat["GPSLongitudeRef"].(string)
	num2 := dat["GPSLongitude"].([]interface{})

	//*** Latitude
	Snum := fmt.Sprintf("%s", num)
	Tnum1 := TrimPrefix(Snum)
	F.Lat = TrimSuffix(Tnum1, "]")

	// *** Longitude
	Snum2 := fmt.Sprintf("%s", num2)
	Tnum2 := TrimPrefix(Snum2)
	F.Long = TrimSuffix(Tnum2, "]")

	return F, nil
}
