package main

import (
	"flag"
	"fmt"
	"github.com/kurtcoke/goexifgps"
	"github.com/rwcarlsen/goexif/exif"
	"io"
	"os"
)

func main() {
	ImageFileName := flag.String("image", "", "Dont use pngs please.")
	flag.Parse()
	var R io.Reader

	if len(*ImageFileName) > 0 {

		Rrr, err := os.Open(*ImageFileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", "Cant open file.")
		}
		R = Rrr
	}
	if len(*ImageFileName) == 0 {
		R = os.Stdin
	}

	ExifData, err := exif.Decode(R)
	if err == io.EOF {
		fmt.Fprintf(os.Stderr, "%s\n", "Error decoding file.")
	}
	if err != io.EOF {
		Location, err := goexifgps.DecodeGPS(ExifData)
		if err != io.EOF {
			fmt.Println(Location)
		}
		if err == io.EOF {
			fmt.Fprintf(os.Stderr, "%s\n", "Contains no GPS exif data.")
		}
	}

}
