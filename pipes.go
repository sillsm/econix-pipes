package pipes

import (
	"io/ioutil"
	"syscall"
        "os"
)

// By default WriteToPipe blocks execution until there's a reader.
func WriteToPipe(filename string, msg []byte) error {
	return ioutil.WriteFile(filename, msg, 0644)
}

// By default ReadFromPipe blocks execution until there's a writer.
func ReadFromPipe(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Make a pipe.
func MakeAPipe(name string) {
	syscall.Mknod(name, syscall.S_IFIFO|0666, 0)
}

// Close a pipe.
func ClosePipe(name string) error {
        return os.Remove(name)
}

// Write to Pipe, Block until content available, Read.
// Insert a callback thing here.
func ReadBlockWrite(f string) (chan []byte, error) {
	out := make(chan []byte)
	go func() {
		request, err := ReadFromPipe(f)
		if err != nil {
			// Do nothing for the moment.
		}
		// Communicate the request to listeners.
		out <- request
		response := <-out
		// Block until other goroutines have created the response.
		err = WriteToPipe(f, response)
		if err != nil {
			// Do nothing for the moment
		}
	}()
	return out, nil
}

// Write to Pipe, Block until content available, Read.
// Insert a callback thing here.
func WriteBlockRead(f string) (chan []byte, error) {
	out := make(chan []byte)
	go func() {
		// Writes to pipe when a signal is available on the channel.
		err := WriteToPipe(f, <-out)
		if err != nil {
			// Do nothing for the moment.
		}
		// Blocks until response available.
		response, err := ReadFromPipe(f)
		if err != nil {
			// Do nothing for the moment.
		}
		// Put the response on the channel.
		out <- response
	}()
	return out, nil
}

