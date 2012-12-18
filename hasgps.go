package GoExifGPS

import (
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
	"io"
	"log"
	"os"
)

type Exif struct {
	tif *tiff.Tiff

	main map[FieldName]*tiff.Tag
}
type FieldName string

func ContainsGPS(fname string) bool {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	f.Close()
	if err != nil {
		return false
		//log.Fatal(err) log.Fatal makes the program stop so we don't need this
		// coz we won't be able to get false ever.
	}
	_, Has := x.Get("GPSLatitude")
	return Has != io.EOF

}

// Remember to add error checking here
func ContainsGPSFromStdin() bool {
	r := os.Stdin
	x, err := exif.Decode(f)
	if err != nil {
		return false
	}
	_, Has := x.Get("GPSLatitude")
	return Has != io.EOF
}
