package pipes

import (
  "syscall"
  "io/ioutil"
)

// By default WriteToPipe blocks execution until there's a reader.
func WriteToPipe(filename string, msg []byte) error{
  return ioutil.WriteFile(filename, msg, 0644)
}

// By default ReadFromPipe blocks execution until there's a writer.
func ReadFromPipe(filename string)([]byte, error){
  return ioutil.ReadFile(filename)
}

// Write to Pipe, Block until content available, Read.
func WriteBlockRead(filename string, msg []byte, resp chan []byte)(error){
  err := WriteToPipe(filename, msg)
  if err != nil{
    return err
  }
  response, err := ReadFromPipe(filename)
  if err != nil{
    return err
  }
  resp <- response
  return nil
}

// Write to Pipe, Block until content available, Read.
// Insert a callback thing here.
func ReadBlockWrite(filename string, com chan []byte)(error){
  request, err := ReadFromPipe(filename)
  if err != nil{
    return err
  }
  // Communicate the request to listeners.
  com <- request
  // Block until other goroutines have created the response.
  msg := <- com
  err = WriteToPipe(filename, msg)
  if err != nil{
    return err
  }
  return nil
}


// Make a pipe.
func MakeAPipe(name string){
  syscall.Mknod(name, syscall.S_IFIFO|0666, 0)
}
