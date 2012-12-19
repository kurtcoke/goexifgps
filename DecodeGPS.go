package GoExifGPS

import (
	"github.com/kurtcc/GoExifGPS"
)

/* Assume we know the file has gps exif data. Use hasgps.go to check for it, has two functions
for doing that.  We don't want to run through this function here completely only to realise we did all that processing just to get an error "Does not contain gps data. */
func DecodeGPS(filename string) (string, error) {
	var Location string
	PointToPic, err := GoExifGPS.OpenClose(filename)
	if err != nil {
		return Location, err
	}
	if err == nil {
		LatRef, Lat, LongRef, LongD := GoExifGPS.OpenParseJson(PointToPic)
		Latitude := GoExifGPS.FormatGPS(Lat)
		Longitude := GoExifGPS.FormatGPS(LongD)
		Latitude = GoExifGPS.RefFormat(LatRef, Latitude)
		Longitude = GoExifGPS.RefFormat(LongRef, Longitude)
		Location = GoExifGPS.MapFriendly(Latitude, Longitude)
		return Location, err
	}
	return Location, err
}
