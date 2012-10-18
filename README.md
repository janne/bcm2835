Go language package to control the Broadcom BCM 2835 as used in the Raspberry
Pi. Builds on Mike McCauley's C lib with the same name.

The package needs Golang 1.1 since the stable 1.0 release doesn't provide a
stable version of cgo. Needs to be run as sudo.

Example:

```go
package main

import (
  "bcm2835"
  "time"
  "fmt"
)

func main() {
  err := bcm2835.Init() // Initialize the library
  if err != nil {
    fmt.Println(err)
    return
  }
  defer bcm2835.Close() // Run close when returning
  bcm2835.GpioFsel(bcm2835.Pin11, bcm2835.Output) // Set pin 11 to output

  for { // Loop forever
    bcm2835.GpioSet(bcm2835.Pin11) // Set pin 11 high
    time.Sleep(500 * time.Millisecond)
    bcm2835.GpioClr(bcm2835.Pin11) // Set pin 11 low
    time.Sleep(500 * time.Millisecond)
  }
}
```
