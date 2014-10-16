package pipes

import (
        "strconv"
	"testing"
)

func TestReadBlockWrite(t *testing.T) {
	MakeAPipe("First")
        defer ClosePipe("First")
	go WriteToPipe("First", []byte("First request"))

	msg, _ := ReadBlockWrite("First")
	req := <-msg
	if string(req) != "First request" {
		t.Errorf("Want %v, got %v ", "First request", string(req))
	}
	msg <- []byte("Response")
	resp, err := ReadFromPipe("First")
	if err != nil {
		t.Errorf("ReadFromPipe: %v", err)
	}
	if string(resp) != "Response" {
		t.Errorf("Got %v, want %v ", string(resp), "Response")
	}
}

func TestWriteBlockRead(t *testing.T) {
	MakeAPipe("Second")
        defer ClosePipe("Second")
	msg, _ := WriteBlockRead("Second")
	msg <- []byte("1. Request")
	req, err := ReadFromPipe("Second")
	if err != nil {
		t.Errorf("ReadFromPipe: %v", err)
	}
	if string(req) != "1. Request" {
		t.Errorf("Got %v, want %v", string(req), "1. Request")
	}

	go WriteToPipe("Second", []byte("2. Response"))
	resp := <-msg
	if string(resp) != "2. Response" {
		t.Errorf("Got %v, want %v", string(resp), "2. Response")
	}
	//msg
}


func TestStressTestWriteAndReadBlock(t *testing.T) {
      limit := 1000 
      requestNum := 10000
      fd := make (chan string, 40000)
      go func () {
        sem := make(chan bool, limit)
        for i := 0; i < requestNum; i++ {
          go func(i int) {
          path := strconv.Itoa(i)
          MakeAPipe(path)
             
          // Scenario: interface got http request, wrote it to pipe.
             sem <- true // Push Semaphore.
             req, _ := WriteBlockRead(path)
             req <- []byte(path + "msg")
             fd <- path
             msg := <- req
             if string(msg) != "Final response" {
               t.Errorf("Got %v, want %v", string(msg), "Final response")
             }
             <- sem      // Pop Semaphore.
          ClosePipe(path)
         }(i)
        }
      }()

     for i := 0; i < requestNum; i++ {
        path := <- fd // Pull a descriptor from written queue.
        req, _  := ReadBlockWrite(path)
        msg := <- req
        if string(msg) != path + "msg" {
           t.Errorf("Got %v, want %v", string(msg), path + "msg")
        }         
        req <- []byte("Final response")
      }

}


//TODO:
//1. Handle errors properly by a struct of stuff
//2. Enable timeouts using the select idiom
//3. Cleanup pipes after you're done by creating a close function.
