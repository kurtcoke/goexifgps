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

To install use:  (Make sure you have git when using go get)

$ sudo apt-get install git

$ sudo go get github.com/rwcarlsen/goexif/tiff

# Exif package depends on the tiff package

# Then go:

$ sudo go get github.com/rwcarlsen/goexif/exif


Why did I write GoExifGPS?
---------------------------

I'm working on a project to scan social media sites.
https://github.com/kurtcc/ExifCrawl (Still very messy.)
For now it will be crawling 4chan's /b. If the crawler finds photo's with GEO/GPS data, then
it saves the photo and produces a document with the photo's name, context (Comments etc.)
and the geo data in a form that can be entered into google maps or arcGIS.

I had some inspiration from http://regex.info/exif.cgi


How to use GoExifGPS?
---------------------
# Make sure you have git. Tried it on a machine that did not have git so go get didn't work.

$ sudo apt-get install git

$ sudo go get github.com/kurtcc/GoExifGPS  
# This should work but you can make sure by navigating to the path something like
$ cd /usr/lib/go/src/pkg/github/GoExifGPS
# and do
$ sudo go install

# Make sure it works
$ wget http://pastebin.com/raw.php?i=1YcxB25P
$ cp raw.php\?i\=1YcxB25P try.go
$ go run try.go

# Make sure you already have goexif installed

Look at try.go at how it should be used in terms of what your code should look like.
Remember to start your code with:
import "github.com/kurtcc/GoExifGPS" 

# Usage examples:

hasgps.go
-----------------
<code>
// Usage: (Example)
package main

import (
"github.com/kurtcc/GoExifGPS"
"fmt"
)

func main() {
        fmt.Println("Lets see if this works.")
         Foo := ContainsGPS("WTgX4.jpg")

         if Foo == false {
        fmt.Println("Image does not contain geo data.")
        }
        if Foo == true {
                fmt.Println("Image contains geo data!")
}
}
</code>
