package main

import (

"github.com/kurtcc/GoExifGPS"
"fmt"


)

func main() {

LatRef, Lat, LongRef,Longd := GoExifGPS.OpenParseJson("_JEF018993_sm.jpg") 
//Returns (string,string,string,string)
// Use in GoExifGPS/parse.go
// TrimSuffix and TrimPrefix are used in a OpenParseJson (Closure)
// You can still use them for something else on their own.

Latitude := GoExifGPS.FormatGPS(Lat)
Longitude := GoExifGPS.FormatGPS(Longd)
// strings are formatted correctly into float32

// Take Ref values and make lat and longitude either + or - based on ref values

Latitude2 := GoExifGPS.RefFormat(LatRef, Latitude) // Might have to change left hand var name
Longitude2 := GoExifGPS.RefFormat(LongRef, Longitude)

location := GoExifGPS.MapFriendly(Latitude2, Longitude2)

fmt.Println(location)


}


