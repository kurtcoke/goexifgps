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
# Make sure you have git. Tried it on a machine that did not have git so go get didn't work.

$ sudo apt-get install git

$ go get github.com/kurtcc/GoExifGPS  
# Make sure your Gopath is set you can check if it is set by running 
$ go env




# Make sure you already have goexif installed

Look at try.go at how it should be used in terms of what your code should look like.
Remember to start your code with:
import "github.com/kurtcc/goexifgps" 

# Usage examples:



</pre>

Using GoExifGPS:
---------------

There will now be two ways to use it. Either by piping an image to stdin:
(This is just a stupid example , there is no error checking etc.
<pre>
$ cat _JEF019769_sm.jpg | try.go
Output: 37.411995,-121.9953
</pre>

The code for try.go looks like this:
<pre>
package main

import (
    "fmt"
    "github.com/kurtcc/goexifgps"
    "github.com/rwcarlsen/goexif/exif"
    "io"
    "os"
)

func main() {
    var R io.Reader
    R = os.Stdin
    ExifData, err := exif.Decode(R)
    if err == io.EOF {
        fmt.Println("Error decoding exif, or now exif in file.")
    }   
    if err != io.EOF {
        Location, err2 := goexifgps.DecodeGPS(ExifData)
        if err2 != io.EOF {
            fmt.Println(Location)
        }   
        if err2 == io.EOF {
            fmt.Println("No GPS data in Exif!")

        }   
    }   

}

</pre>

You can also use this method

<pre>
$ go run cltool.go -image=_JEF019769_sm.jpg 
37.411995,-121.9953
</pre>

If you would like to use it this way,your  code should look like this:
<pre>



</pre>


