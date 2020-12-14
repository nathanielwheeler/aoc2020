package main

import (
	"log"
	"os"
)

func main() {
  // start server
	s := server{
		logger: log.New(os.Stdout, "aoc2020: ", log.Lshortfile),
  }
  
  // get args
  day := os.Args[1]

  switch day {
  case "1.1":
    s.day1p1()
  case "1.2":
    s.day1p2()
  default:
    s.logMsg("Invalid arg input.  Format: '<n>.<1|2>")
  }
}

type server struct {
	logger *log.Logger
}