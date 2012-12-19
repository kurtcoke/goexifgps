package GoExifGPS
// Author kurtcc on github.com
// use like so Location, err := GoExifGPS.DecodeGPS("imagename.jpg)
// if err != nil { panic(err) }
// fmt.Println(Location) 

/* Assume we know the file has gps exif data. Use hasgps.go to check for it, has two functions
for doing that.  We don't want to run through this function here completely only to realise we did all that processing just to get an error "Does not contain gps data. */
func DecodeGPS(filename string) (string, error) {
	var Location string
	PointToPic, err := OpenClose(filename)
	if err != nil {
		return Location, err
	}
	if err == nil {
		LatRef, Lat, LongRef, LongD := OpenParseJson(PointToPic)
		Latitude := FormatGPS(Lat)
		Longitude := FormatGPS(LongD)
		Latitude = RefFormat(LatRef, Latitude)
		Longitude = RefFormat(LongRef, Longitude)
		Location = MapFriendly(Latitude, Longitude)
		return Location, err
	}
	return Location, err
}
