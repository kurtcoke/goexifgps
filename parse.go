package goexifgps

// Author : kurtcc on github
// This will be called parse.go
import (
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
	"os"
)

type Exif struct {
	tif *tiff.Tiff

	main map[FieldName]*tiff.Tag
}
type FieldName string

type GeoFields struct {
	LatRef, LongRef string
	Lat, Long       float32
}

//Opens using filename and does exif.Decode
// filename will be known if it is passed as flag via commandline.
func OpenClose(filename string) (*exif.Exif, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	ExifData, err2 := exif.Decode(f)
	f.Close()
	if err2 != nil {
		return nil, err2
	}
	return ExifData, nil

}

// Returns pointer to struct GeoFields 
// can easily be accessed E.Lat, E.Long etc       
func GetGPS(E *exif.Exif) (*GeoFields, error) {

	F := new(GeoFields)
	LatVal, err := E.Get("GPSLatitude")
	if err != nil {
		return nil, err
	}
	F.Lat = FormatGPS(LatVal)
	LongVal, err := E.Get("GPSLongitude")
	if err != nil {
		return nil, err
	}
	F.Long = FormatGPS(LongVal)

	LatRefVal, err := E.Get("GPSLatitudeRef") //Lat and LatRef
	LongRefVal, err := E.Get("GPSLongitudeRef")

	F.LatRef = LatRefVal.StringVal()
	F.LongRef = LongRefVal.StringVal()

	// *** Longitude

	return F, nil
}
