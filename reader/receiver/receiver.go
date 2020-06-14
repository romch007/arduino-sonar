package receiver

import ( 
  "fmt"
  "log"
  "strings"
  "strconv"

	"github.com/tarm/serial"
)

type Record struct {
  Angle, Distance int
}

func parseRecord(rawStr string) (record Record) {
  strs := strings.Split(rawStr, " ")

  angle, _ := strconv.Atoi(strs[0])
  distance, _ := strconv.Atoi(strs[1])

  record = Record{Angle: angle, Distance: distance}

  return
}

func StartReceiver(data chan Record) {
	fmt.Println("Start receiving...")
	c := &serial.Config{Name: "COM5", Baud: 9600}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

  for {
    buf := make([]byte, 512)
    n, err := s.Read(buf)

    if err != nil {
      log.Fatal(err)
    }

    data <- parseRecord(string(buf)) 
  }
}
