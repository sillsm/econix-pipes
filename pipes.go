package pipes

import (
  "syscall"
  "io/ioutil"
)

//By default WriteToPipe blocks execution until there's a reader.
func WriteToPipe(filename string, msg []byte) error{
  return ioutil.WriteFile(filename, msg, 0644)
}

//By default ReadFromPipe blocks execution until there's a writer.
func ReadFromPipe(filename string)([]byte, error){
  return ioutil.ReadFile(filename)
}

func MakeAPipe(name string){
  syscall.Mknod(name, syscall.S_IFIFO|0666, 0)
}
