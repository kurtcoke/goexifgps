goexifgps
=========

Just a small library to get GPS values extracted from jpeg's using rwcarlsen's goexif 
library into a form readable by google maps and similiar mapping tools e.g arcGIS.

cmd.go is a commandline tool. You can use a bash script to scan large amounts of photo's with the command line tool and just write stdout to a text file. 

goexif:
-------

Go exif is by rwcarlsen on github here is the link to goexif on his page:
https://github.com/rwcarlsen/goexif

The docs for goexif are here:
http://go.pkgdoc.org/github.com/rwcarlsen/goexif/exif


To install: 


$ go get github.com/rwcarlsen/goexif/tiff

$ go get github.com/rwcarlsen/goexif/exif


$ go get github.com/kurtcc/goexifgps


Using cmd.go:
---------------


<pre>
$ cat _JEF019769_sm.jpg | try.go
Output: 37.411995,-121.9953
</pre>



