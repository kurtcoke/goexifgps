package parse

// This will be called parse.go
import (
	"encoding/json"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"strings"

//This is how to print a given field
)

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func TrimPrefix(s string) string {
	s = s[1:]
	return s
}

// Use like this
//LatRef, Lat, LongRef,Longd := OpenParseJson("_JEF018993_sm.jpg") 
func OpenParseJson(filename string) (string, string, string, string) {
	// I want this to return all four values each as a string.	
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		panic(err)
	}
	if x == nil {
		fmt.Println(" ")
	}

	b, err := x.MarshalJSON()
	if err != nil {
		panic(err) //Format die output properly
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(b, &dat); err != nil {
		panic(err)
	}

	numR := dat["GPSLatitudeRef"].(string) //Lat and LatRef
	num := dat["GPSLatitude"].([]interface{})

	num2R := dat["GPSLongitudeRef"].(string)
	num2 := dat["GPSLongitude"].([]interface{})

	//*** Latitude
	Snum := fmt.Sprintf("%s", num)
	Tnum1 := TrimPrefix(Snum)
	Tnum := TrimSuffix(Tnum1, "]")

	// *** Longitude
	Snum2 := fmt.Sprintf("%s", num2)
	Tnum2 := TrimPrefix(Snum2)
	Tnum_2 := TrimSuffix(Tnum2, "]")

	return numR, Tnum, num2R, Tnum_2
}
