package main

import (
	"fmt"
	"io"
	"github.com/kurtcc/goexifgps"
)

func main() {
	//FalseImage := "testdata/WTgX4.jpg"
	// Does not contain GPS data.
	// First will return string, latter will return error.
	//TrueImage := "_JEF018993_sm.jpg"
    Image := "WTgX4.jpg"
	//Image := "_JEF018993_sm.jpg" 
	ExifData, err := goexifgps.OpenClose(Image)
	if err == io.EOF {
		fmt.Println("No exif data in file!")
		}
	if err != io.EOF {



	//F := goexifgps.GetGPS(ExifData)
	Location, err2 := goexifgps.DecodeGPS(ExifData)
    if err2 == io.EOF { 
		fmt.Println("No GPS data in Exif!")
	}
	if err2 != io.EOF {

	fmt.Println(Location)
                      }
}              
}
