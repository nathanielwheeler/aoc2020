package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
Puzzle: count number of valid passports, with every field needed _except_ for 'cid'
- Fields are separated by spaces or newlines
- Passports are separated by blank lines

Fields
    byr (Birth Year)
    iyr (Issue Year)
    eyr (Expiration Year)
    hgt (Height)
    hcl (Hair Color)
    ecl (Eye Color)
    pid (Passport ID)
    cid (Country ID)        // optional

*/

func (s *server) day4p1() {
	data, err := ioutil.ReadFile("input/day4.txt")
	if err != nil {
		s.logErr("read file failed", err)
	}

	xs := strings.Split(string(data), `

`)

  fmt.Println("total passports: ", len(xs))
  
  reqFields := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:",}

  validPassports := 0

  for _, passport := range xs {
    validPassports++
    for _, field := range reqFields {
      if !strings.Contains(passport, field) {
        validPassports--
        break
      }
    }
  }

  fmt.Println("valid passports: ", validPassports) // 226
}
