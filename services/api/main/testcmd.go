package main

import (

  "os/exec"
  "fmt"
  "io"
  "bytes"
)
var DLPath = "/home/ubuntu/dev/src/deeplearning/word2vec-read-only/"
var cmd = exec.Command(DLPath + "distance", DLPath + "freebase-vectors-skipgram1000-en.bin" )



func main(){

  //cmd:= exec.Command("echo" , "testing")
  //cmd.Start()

  err := cmd.Start()
  
  fmt.Println(err)
  stdout , err := cmd.StdoutPipe()

  /*

  if err == nil {
    fmt.Println(out)
  } 
  
  */
   

   

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely

    go func() {

        var buf bytes.Buffer
        io.Copy( &buf , stdout )
        outC <- buf.String()
        
    }()

        //out := <-outC
        fmt.Println(outC)
    

}