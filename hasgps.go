package GoExifGPS

import (
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
	"io"
)

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

