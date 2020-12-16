package main

import (
	"errors"
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

	// separate into slice of strings representing passports
	xs := strings.Split(string(data), "\n\n")
	fmt.Println("total passports: ", len(xs))

	// further separate each passport into a map[string]string
	var xDict []map[string]string
	for _, pass := range xs {

		// Normalize fields into pp
		var pp []string
		// split by newlines
		ss := strings.Split(pass, "\n")
		for _, str := range ss {
			// if a string contains a space, split that, appending to pp regardless
			if strings.Contains(str, " ") {
				sss := strings.Split(str, "\n")
				pp = append(pp, sss...)
			} else {
				pp = append(pp, str)
			}
		}

		dict := make(map[string]string)
		// go through pp and split by ':', adding this key value pair to dict
		for _, str := range pp {
			kv := strings.Split(str, ":")
			dict[kv[0]] = kv[1]
		}

		// validate presence of required fields
		reqFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		validDict := true
		for _, field := range reqFields {
			_, ok := dict[field]
			if !ok {
				validDict = false
				break
			}
		}

		if validDict {
			xDict = append(xDict, dict)
		}
	}
	fmt.Println("xDict: ", len(xDict)) // this should be 226

	// parse xDict into a slice of passports
	var xPass []passport
	for _, dict := range xDict {
		pass := passport{
			birthYear: dict["byr"],
			issueYear: dict["iyr"],
			expYear:   dict["eyr"],
			height:    dict["hgt"],
			hairColor: dict["hcl"],
			eyeColor:  dict["ecl"],
			passID:    dict["pid"],
		}
		xPass = append(xPass, pass)
	}

	fmt.Println("xPass: ", len(xPass)) // this should also be 226

	// Run validations
	validPassports := 0
	for _, pass := range xPass {
		err := s.runPassportValFns(&pass,
			s.passByr,
			s.passIyr,
			s.passEyr,
			s.passHgt,
			s.passHcl,
			s.passEcl,
			s.passPid,
		)
		if err != nil {
			continue
		}
		validPassports++
	}

	fmt.Println("valid passports: ", validPassports) // ?
}

type passport struct {
	birthYear string
	issueYear string
	expYear   string
	height    string
	hairColor string
	eyeColor  string
	passID    string
}

type passportValFn func(*passport) error

func (s *server) runPassportValFns(pass *passport, fns ...passportValFn) error {
	for _, fn := range fns {
		if err := fn(pass); err != nil {
			return err
		}
	}
	return nil
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func (s *server) passByr(pass *passport) error {
	byr, err := strconv.Atoi(pass.birthYear)
	if err != nil {
		return err
	}
	if byr <= 1920 && byr <= 2002 {
		return errors.New("birth year is too high or low")
	}
	return nil
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func (s *server) passIyr(pass *passport) error {
	iyr, err := strconv.Atoi(pass.issueYear)
	if err != nil {
		return err
	}

	if iyr <= 2010 && iyr <= 2020 {
		return errors.New("issue year is too high or low")
	}
	return nil
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func (s *server) passEyr(pass *passport) error {
	eyr, err := strconv.Atoi(pass.issueYear)
	if err != nil {
		return err
	}

	if eyr <= 2020 && eyr <= 2030 {
		return errors.New("expiration year is too high or low")
	}
	return nil
}

// hgt (Height) - a number followed by either cm or in:
//     If cm, the number must be at least 150 and at most 193.
//     If in, the number must be at least 59 and at most 76.
func (s *server) passHgt(pass *passport) error {
	hgt := pass.height
	errInvalidHgt := errors.New("height is too big or small")

	if strings.HasSuffix(hgt, "cm") {
		hgt = strings.TrimSuffix(hgt, "cm")
		cm, err := strconv.Atoi(hgt)
		if err != nil {
			return err
		}
		if cm >= 150 && cm <= 193 {
			return errInvalidHgt
		}
	} else if strings.HasSuffix(hgt, "in") {
		hgt = strings.TrimSuffix(hgt, "in")
		in, err := strconv.Atoi(hgt)
		if err != nil {
			return err
		}
		if in >= 59 && in <= 76 {
			return errInvalidHgt
		}
	} else {
		return errors.New("invalid height data")
	}
	return nil
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func (s *server) passHcl(pass *passport) error {
	var colorHex = regexp.MustCompile(`^[0-9a-f]{6}$`)
	hcl := pass.hairColor
	if colorHex.MatchString(hcl) {
		return errors.New("invalid hair color value")
	}
	return nil
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func (s *server) passEcl(pass *passport) error {
	ecl := pass.eyeColor
	errInvalidEcl := errors.New("invalid eye color")
	if len(ecl) != 3 {
		return errInvalidEcl
	}
	switch ecl {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
		return nil
	}
	// if none of the above are true
	return errInvalidEcl
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func (s *server) passPid(pass *passport) error {
	nineDigits := regexp.MustCompile(`^[0-9]{9}$`) // including leading zeroes
	pid := pass.passID
	if nineDigits.MatchString(pid) {
		return errors.New("invalid passport id")
	}
	return nil
}
