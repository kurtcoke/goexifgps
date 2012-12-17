TODO:
-----

Write generate documentation for the code.

Later on change float32 to float64 to see if it can improve precision, for now I think it 
is pretty precise. I compared it to the tool on regex.info and I seem to land on the same 
house using my tool and using the tool on regex.ingo

This is how to open file piped to stdin
(As in $ cat _JEF018993_sm.jpg | go run stdin.go)

<pre>
package main

import (
  "os"
  "log"
  "fmt"

  "github.com/rwcarlsen/goexif/exif"
)

func main() {
  

  

  x, err := exif.Decode(os.Stdin)
  //f.Close()
  if err != nil {
    log.Fatal(err)
  }

  camModel, _ := x.Get("Model")
  date, _ := x.Get("DateTimeOriginal")
  fmt.Println(camModel.StringVal())
  fmt.Println(date.StringVal())

  focal, _ := x.Get("FocalLength")
  numer, denom := focal.Rat2(0) // retrieve first (only) rat. value
  fmt.Printf("%v/%v", numer, denom)
}
//   x, err := exif.Decode(os.Stdin)
// run like so 
// $ cat _JEF018993_sm.jpg | go run stdin.go
</pre>
