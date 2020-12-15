package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

	fmt.Println("total passports: ", len(xs)) // 280

	reqFields := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}

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

func (s *server) day4p2() {
	data, err := ioutil.ReadFile("input/day4.txt")
	if err != nil {
		s.logErr("read file failed", err)
	}

  xs := strings.Split(string(data), `

`)

	fmt.Println("total passports: ", len(xs))

	var xxs [][]string
	for _, passport := range xs {
		var pp []string
		ss := strings.Split(passport, " ")
		for _, str := range ss {
			if strings.Contains(str, `
`) {
				sss := strings.Split(str, "\n")
				pp = append(pp, sss...)
			} else {
				pp = append(pp, str)
			}
		}
		xxs = append(xxs, pp)
	}

	validPassports := 0
	reqFields := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
	/* Validations

	 */
	var (
    colorHex = regexp.MustCompile(`^[0-9a-f]{6}$`)
    nineDigits = regexp.MustCompile(`^[0-9]{9}$`)// including leading zeroes
	)

	for _, passport := range xxs {
		isValidPassport := true
		for _, req := range reqFields {
			isValidField := false
			for _, field := range passport {
				if strings.HasPrefix(field, req) {
					// If I get here, that means that it has the required field that I am looking at.  That means it still needs to pass validation.

					// trim prefix so that i can work with just the value
					field = strings.TrimPrefix(field, req)

					switch req {
					// byr (Birth Year) - four digits; at least 1920 and at most 2002.
					case reqFields[0]:
            byr, err := strconv.Atoi(field)
            if err != nil {
              break
            }
            if byr <= 1920 && byr <= 2002 {
              isValidField = true
              break
            }
          // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
          case reqFields[1]:
            iyr, err := strconv.Atoi(field)
            if err != nil {
              break
            }
            if iyr <= 2010 && iyr <= 2020 {
              isValidField = true
              break
            }
          // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
          case reqFields[2]:
            eyr, err := strconv.Atoi(field)
            if err != nil {
              break
            }
            if eyr <= 2020 && eyr <= 2030 {
              isValidField = true
              break
            }
          // hgt (Height) - a number followed by either cm or in:
        case reqFields[3]:
          //     If cm, the number must be at least 150 and at most 193.
          //     If in, the number must be at least 59 and at most 76.
            if strings.HasSuffix(field, "cm") {
              field = strings.TrimSuffix(field, "cm")
              hgt, err := strconv.Atoi(field)
              if err != nil {
                break
              }
              if hgt >= 150 && hgt <= 193 {
                isValidField = true
                break
              }
            } else if strings.HasSuffix(field, "in") {
              field = strings.TrimSuffix(field, "in")
              hgt, err := strconv.Atoi(field)
              if err != nil {
                break
              }
              if hgt >= 59 && hgt <= 76 {
                isValidField = true
                break
              }
            } else {
              break
            }
          // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
          case reqFields[4]:
            if colorHex.MatchString(field) {
              isValidField = true
              break
            }
          // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
          case reqFields[5]:
            if len(field) != 3 {
              break
            }
            switch field {
            case "amb":
            case "blu":
            case "brn":
            case "gry":
            case "grn":
            case "hzl":
            case "oth":
              isValidField = true
            }
          // pid (Passport ID) - a nine-digit number, including leading zeroes.
          case reqFields[6]:
            if nineDigits.MatchString(field) {
              isValidField = true
            }
					}
				}
			}
			if !isValidField {
				isValidPassport = false
				break
			}
		}
		if isValidPassport {
			validPassports++
		}
	}

	fmt.Println("valid passports: ", validPassports) // ?
}
