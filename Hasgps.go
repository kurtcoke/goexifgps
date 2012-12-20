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
	var Contains bool
	Contains, err = ContainsGPSFromStdin(f)
	/* if err2 == nil {
	return Contains, nil
	} else {
		return false ,err2
	}
	// Ok fine this works , later make ContainsGPSFromStdin return errors too
	*/
	return Contains, err
}

// Remember to add error checking here
func ContainsGPSFromStdin(R io.Reader) (bool, error) {

	x, err := exif.Decode(R)
	if err != nil {
		return false, err
	}
	_, err = x.Get("GPSLatitude") //returns (*tiff.Tag, error)

	return err != io.EOF, err

}
