package goexifgps

import "github.com/rwcarlsen/goexif/exif"

// Author kurtcc on github.com

func DecodeGPS(ExifData *exif.Exif) (string, error) {
	// (string,error)
	F, err := GetGPS(ExifData)
	if err != nil {
		return "", err //maybe log(err)
	}

	Latitude := RefFormat(F.LatRef, F.Lat)
	Longitude := RefFormat(F.LongRef, F.Long)
	Location := MapFriendly(Latitude, Longitude)
	return Location, nil
}
