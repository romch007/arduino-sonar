package receiver

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tarm/serial"
)

type Record struct {
	Angle, Distance int
}

func isEndingMessage(rawString string) bool {

	return strings.Compare(strings.TrimSuffix(rawString, "\n"), "end") == 0
}

func parseRecord(rawRecord string) (record *Record) {

	strs := strings.Split(rawRecord, ",")

	angle, _ := strconv.Atoi(strs[0])
	distance, _ := strconv.Atoi(strings.TrimSuffix(strs[1], "\n"))

	record = &Record{Angle: angle, Distance: distance}

	return
}

func StartReceiver(recordsChan chan<- *Record) {

	fmt.Println("Start receiving...")
	c := &serial.Config{Name: "COM5", Baud: 9600}
	s, err := serial.OpenPort(c)
	defer s.Close()

	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(s)
		received, err := reader.ReadBytes('\x0a')

		if err != nil {
			panic(err)
		}

		toString := string(received)

		if isEndingMessage(toString) {
			fmt.Println("End signal received")
			close(recordsChan)
			break
		} else {
			recordsChan <- parseRecord(toString)
		}
	}
}
