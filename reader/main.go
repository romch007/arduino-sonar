package main

import (
	"sync"

	"gitlab.com/romch007/sonar/reader/graphic"
	"gitlab.com/romch007/sonar/reader/receiver"
)

func main() {
	dataChan := make(chan *receiver.Record)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		receiver.StartReceiver(dataChan)
	}()
	go func() {
		defer wg.Done()
		graphic.StartGraphic(dataChan)
	}()

	wg.Wait()
}
