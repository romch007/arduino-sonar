package graphic

import (
	"fmt"

	"gitlab.com/romch007/sonar/reader/receiver"
)

func StartGraphic(recordsChan <-chan *receiver.Record) {
	fmt.Println("Starting graphics...")
	go serveHTML()

	for {
		incoming := <-recordsChan
		// fmt.Println("Receive", incoming)
		sendToClient(incoming)
	}
}
