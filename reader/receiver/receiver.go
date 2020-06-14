package receiver

import ( 
  "fmt"
  "log"
  "strings"
  "strconv"
  "bufio"

	"github.com/tarm/serial"
)

type Record struct {
  Angle, Distance int
}

func parseRecord(rawRecord string) (record Record) {
  strs := strings.Split(rawRecord, ",")

  fmt.Println("strs:", strs)

  angle, _ := strconv.Atoi(strs[0])
  distance, _ := strconv.Atoi(strs[1])

  fmt.Println("angle:", angle)
  fmt.Println("distance:", distance)

  record = Record{Angle: angle, Distance: distance}

  return
}

func StartReceiver(data chan Record) {
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

    data <- parseRecord(string(received))
  }
}
