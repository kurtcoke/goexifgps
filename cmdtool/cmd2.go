package main

import (
	"flag"
	"fmt"
	"github.com/kurtcc/goexifgps"
	"github.com/rwcarlsen/goexif/exif"
	"io"
	"os"
)

func main() {
 	Rr := flag.String("image", "", "Dont use pngs please.")    
	flag.Parse()
	What := flag.NFlag()
	var R io.Reader
	
	switch What {
	case 0:
		R = os.Stdin
		break
	case 1:
		
		Rrr, err := os.Open(*Rr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", "Cant open file.")
		}  
		R = Rrr
		break
	}

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
