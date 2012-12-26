package GoExifGPS
// Author kurtcc on github.com

/*
ExifData , err := exif.Decode(r io.Reader) Takes an io.Reader
Location, err := GoExifGPS.DecodeGPS(ExifData)

Or 

ExifData, err := OpenClose(filename)
Location , err2 := GoExifGPS.DecodeGPS(ExifData)
*/

/* Assume we know the file has gps exif data. Use hasgps.go to check for it, has two functions
for doing that.  We don't want to run through this function here completely only to realise we did all that processing just to get an error "Does not contain gps data. */


func DecodeGPS(ExifData *exif.Exif) (string, error) {
		LatRef, Lat, LongRef, LongD := OpenParseJson(ExifData)
		Latitude := FormatGPS(Lat)
		Longitude := FormatGPS(LongD)
		Latitude = RefFormat(LatRef, Latitude)
		Longitude = RefFormat(LongRef, Longitude)
		Location = MapFriendly(Latitude, Longitude)
		return Location, err
}
