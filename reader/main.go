package main

import (
	"fmt"

  "gitlab.com/romch007/sonar/reader/receiver"
)

func main() {
  dataChan := make(chan receiver.Record)
  go receiver.StartReceiver(dataChan);

  for {
    fmt.Println(<-dataChan)
  }
}
