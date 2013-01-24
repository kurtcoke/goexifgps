package main

import (
	"fmt"
	"github.com/kurtcc/goexifgps"
	"github.com/rwcarlsen/goexif/exif"
	"io"
	"os"
)

func main() {
	var R io.Reader
	R = os.Stdin
	ExifData, err := exif.Decode(R)
	if err == io.EOF {
		fmt.Fprintf(os.Stderr, "%s\n", "Error decoding file.")
	}
	if err != io.EOF {
		Location, err2 := goexifgps.DecodeGPS(ExifData)
		if err2 != io.EOF {
			fmt.Println(Location)
		}
		if err2 == io.EOF {
			fmt.Fprintf(os.Stderr, "%s\n", "Contains no GPS exif data.")
		}
	}

}
