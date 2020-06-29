package graphic

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
  "errors"

	"github.com/gorilla/websocket"
	"gitlab.com/romch007/sonar/reader/receiver"
)

var conn *websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		messageType, p, err := conn.ReadMessage()
		fmt.Println(messageType, p, err)
	}
}

func parseRecord(record *receiver.Record) []byte {
	angleStr := strconv.Itoa(record.Angle)
	distanceStr := strconv.Itoa(record.Distance)
	str := distanceStr + " " + angleStr
	return []byte(str)
}

func sendToClient(message *receiver.Record) (err error) {
  if conn == nil {
    err = errors.New("No connection")
    return
  }
	err = conn.WriteMessage(1, parseRecord(message))
	return
}
