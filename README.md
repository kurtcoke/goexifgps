goexifgps
=========

Just a small library to get GPS values extracted from jpeg's using rwcarlsen's goexif library into a form readable by google maps and similiar mapping tools e.g arcGIS.

goexif:
-------

Go exif is by rwcarlsen on github here is the link to goexif on his page:
https://github.com/rwcarlsen/goexif

The docs for goexif are here:
http://go.pkgdoc.org/github.com/rwcarlsen/goexif/exif

NB: Exif package depends on the tiff package.

To install use:  (Make sure you have git when using go get)

$ sudo apt-get install git

$ sudo go get github.com/rwcarlsen/goexif/tiff

# Exif package depends on the tiff package

# Then go:

$ sudo go get github.com/rwcarlsen/goexif/exif


How to use GoExifGPS?
---------------------

$ go get github.com/kurtcc/GoExifGPS  
# Make sure your Gopath is set you can check if it is set by running 



# Usage examples:



</pre>

Using cmd.go:
---------------


<pre>
$ cat _JEF019769_sm.jpg | try.go
Output: 37.411995,-121.9953
</pre>



