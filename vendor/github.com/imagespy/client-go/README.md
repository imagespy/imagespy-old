# client-go

## Usage

```go
package main

import (
  "log"

  "github.com/imagespy/client-go"
)

func main() {
  client := imagespy.NewClientV1()
  spy, err := client.ImageSpy.Get("golang:1.9.1")
  if err != nil {
    log.Fatal(err)
  }

  log.Println(spy.Name)
  log.Println(spy.CurrentImage.Tag)
  log.Println(spy.LatestImage.Tag)
}
```
