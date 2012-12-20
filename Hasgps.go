package GoExifGPS

import (
	"github.com/rwcarlsen/goexif/exif"
	"io"
	"os"
)

// I need this: especially for walking folders.
func ContainsGPS(fname string) (bool, error) {
	f, err := os.Open(fname)
	if err != nil {
		return false, err
	}
	Contains, err2 := ContainsGPSFromStdin(f)
	return Contains, err
	// Ok fine this works , later make ContainsGPSFromStdin return errors too

}

// Remember to add error checking here
func ContainsGPSFromStdin(R io.Reader) (bool, error) {
	x, err := exif.Decode(R)
	if err != nil {
		return false, err
	}
	_,err2 := x.Get("GPSLatitude")   //returns (*tiff.Tag, error)
	if err2 == nil {
		return true, nil
	}
	if err2 != nil {
		return false, err2
}                                                            
}
