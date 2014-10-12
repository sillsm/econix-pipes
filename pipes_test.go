package pipes

import (
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

//TODO:
//1. Handle errors properly by a struct of stuff
//2. Enable timeouts using the select idiom
//3. Cleanup pipes after you're done by creating a close function.
