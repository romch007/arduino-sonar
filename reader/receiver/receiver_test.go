package receiver

import (
  "testing"
  "reflect"
)

func TestParseRecord(t *testing.T) {
  sample := "34 56"
  got := parseRecord(sample)
  result := Record{Angle: 34, Distance: 56}

  if !reflect.DeepEqual(got, result) {
    t.Error("Wrong result")  
  }
}
