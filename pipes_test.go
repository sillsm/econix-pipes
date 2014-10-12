package pipes

import (
	"testing"
)

func TestReadBlockWrite(t *testing.T) {
	MakeAPipe("First")
	go WriteToPipe("First", []byte("First request"))

	var comms = make(chan []byte)
	go ReadBlockWrite("First", comms)
	req := <-comms
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

func TestBetterReadBlockWrite(t *testing.T) {
	MakeAPipe("First")
	go WriteToPipe("First", []byte("First request"))

	msg, _ := BetterReadBlockWrite("First")
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

func TestBetterWriteBlockRead(t *testing.T) {
	MakeAPipe("Second")
	//defer *to implement a close function*
	msg, _ := BetterWriteBlockRead("Second")
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
