package graphic

import (
	"fmt"
	"math"

	"github.com/fogleman/gg"
	"gitlab.com/romch007/sonar/reader/receiver"
)

const (
	Width = 255
	Height
)

func degreeToRadian(angle float64) float64 {
	return angle * math.Pi / 180
}

func convertCoordinates(record *receiver.Record) (x, y float64) {
	x = float64(record.Distance) * math.Cos(degreeToRadian(float64(record.Angle)))
	y = float64(record.Distance) * math.Sin(degreeToRadian(float64(record.Angle)))
	x = Width/2 + x
	y = Height/2 - y
	return
}

func isClosed(ch <-chan *receiver.Record) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func StartGraphic(recordsChan <-chan *receiver.Record) {
	fmt.Println("Starting graphics...")
	dc := gg.NewContext(Width, Height)
	for {
		if isClosed(recordsChan) {
			fmt.Println("Saving image...")
			dc.SavePNG("out.png")
			fmt.Println("Saved")
			break
		} else {
			incoming := <-recordsChan
			//fmt.Println("Receive", incoming)
			if incoming == nil {
				continue
			}
			x, y := convertCoordinates(incoming)
			fmt.Printf("(%f, %f)\n", x, y)

			dc.DrawPoint(x, y, float64(2.0))
			dc.SetRGB255(255, 0, 0)
			dc.Fill()
		}
	}
}
