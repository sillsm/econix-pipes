package pipes

import (
	"testing"
)

func TestServerClient(t *testing.T){
  MakeAPipe("First")
  go WriteToPipe("First", []byte("First msg\n"))
  go WriteToPipe("First", []byte("Second msg\n"))
  go WriteToPipe("First", []byte("Third msg\n"))
  
   
  b, err := ReadFromPipe("First")
  if err != nil {
    t.Errorf("ReadFromPipe(\"First\"): %v ", err)
  }
  
  if string(b) != "First msg\n" {
    t.Errorf("Want %v, got %v ", "First msg\n", string(b))
  }
}
