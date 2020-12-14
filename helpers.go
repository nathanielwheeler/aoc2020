package main

import (
	"bufio"
	"os"
	"strconv"
)

// parseFileToXInt takes in a relative filename and parses each line into a slice of int, performing a panic if there is any error.
func (s *server) parseFileToXInt(filename string) []int {
  filename = "input/" + filename + ".txt"

	data, err := os.Open(filename)
	if err != nil {
		s.panic("Couldn't open file", err)
	}

	scanner := bufio.NewScanner(data)
	var xi []int

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			s.panic("integer conversion failed", err)
		}
		xi = append(xi, line)
	}

	return xi
}

// parseFileToXStr takes in a relative filename and parses each line into a slice of string, performing a panic if there is any error.
func (s *server) parseFileToXStr(filename string) []string {
  filename = "input/" + filename + ".txt"

	data, err := os.Open(filename)
	if err != nil {
		s.panic("Couldn't open file", err)
	}

	scanner := bufio.NewScanner(data)
	var xs []string

	for scanner.Scan() {
		line := scanner.Text()
		xs = append(xs, line)
	}

	return xs
}

// logMsg takes in a message string and outputs it to the server's logger.
func (s *server) logMsg(msg string) {
	s.logger.Println(msg)
}

// logErr takes in a message string and an error and outputs them to the server's logger.
func (s *server) logErr(msg string, err error) {
	s.logger.Printf("%s\n\t%s\n", msg, err)
}

// panic will take in a message and an error, give them to the server's logger, and make a call to panic().
func (s *server) panic(msg string, err error) {
	s.logger.Panicf("%s\n\t%s\n", msg, err)
}
