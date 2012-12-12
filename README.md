GoExifGPS
=========

Just a small library to get GPS values extracted from jpeg's using rwcarlsen's goexif library into a form readable by google maps and similiar mapping tools e.g arcGIS.

Goexif:
-------

Go exif is by rwcarlsen on github here is the link to goexif on his page:
https://github.com/rwcarlsen/goexif
The docs for goexif are here:
http://go.pkgdoc.org/github.com/rwcarlsen/goexif/exif

NB: Exif package depends on the tiff package.

To install use:

go get github.com/rwcarlsen/goexif/tiff

# Exif package depends on the tiff package

# Then go:

go get github.com/rwcarlsen/goexif/exif


Why did I write GoExifGPS?
---------------------------

I'm working on a project to scan social media sites, but for
now it will be crawling 4chan's /b. If the crawler finds photo's with GEO/GPS data, then
it saves the photo and produces a document with the photo's name, context (Comments etc.)
and the geo data in a form that can be entered into google maps or arcGIS.

I had some inspiration from http://regex.info/exif.cgi

