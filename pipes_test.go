package pipes

import (
	"testing"
)

func TestReadBlockWrite(t *testing.T){
  MakeAPipe("First")
  go WriteToPipe("First", []byte("First request"))

  var comms = make(chan []byte)
  go ReadBlockWrite ("First", comms)  
  req := <- comms
  if string(req) != "First request" {
    t.Errorf("Want %v, got %v ", "First request", string(req))
  }
  comms <- []byte("Response")
  resp, err := ReadFromPipe("First")
  if err != nil {
    t.Errorf("ReadFromPipe: %v", err)
  }
  if string(resp) != "Response" {
    t.Errorf("Got %v, want %v ", string(resp), "Response")
  } 
}
