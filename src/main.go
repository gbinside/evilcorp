package main

import (
  "flag"
  "censorship"
  "runtime/pprof"
  "log"
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

  for i:=0; i<100000; i++ {
    censorship.Censorship(args[0])
  }
}
