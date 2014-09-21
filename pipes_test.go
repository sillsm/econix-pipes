package pipes

import (
	"testing"
)

func TestServerClient(t *testing.T){
  MakeAPipe("First")
  
  var response = make(chan []byte)
  var answer = make([]byte)
  
  go WriteBlockRead("First, []byte("First msg"), response)
  go WriteToPipe("First", []byte("Second msg\n"))
  
  answer <- response
  if string(answer) != "Second msg" {
    t.Errorf("Want %v, got %v ", "Second msg", string(answer))
  }
  
  /*
  
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
  */
}
