package main

import (
	"fmt"
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
/* Usage: (Example)
package main

import ( 
"github.com/kurtcc/GoExifGPS"
"fmt"
)

func main() {
	fmt.Println("Lets see if this works.")
	 Foo := ContainsGPS("WTgX4.jpg")

	 if Foo == false {
	fmt.Println("Image does not contain geo data.")
	}
	if Foo == true { 
		fmt.Println("Image contains geo data!")
}
}
 */
