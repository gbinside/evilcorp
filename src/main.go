package main

import (
  "flag"
  "censorship"
  "runtime/pprof"
  "log"
  "fmt"
  "os"
)

func main() {

  f, err := os.Create("main.prof")
  if err != nil {
    log.Fatal(err)
  }
  pprof.StartCPUProfile(f)
  defer pprof.StopCPUProfile()

  flag.Parse()
  args := flag.Args()

  for i:=0; i<10000; i++ {
    censorship.Censorship(args[0])
  }
  fmt.Println(censorship.Censorship(args[0]))

}
