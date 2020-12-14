package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc

Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

How many passwords are valid according to their policies?
*/
func (s *server) day2p1() {
	xs := s.parseFileToXStr("day2")

	var valid int

	for _, str := range xs {
		split := strings.Split(str, " ")
		split[1] = strings.TrimSuffix(split[1], ":")
		pwRune := split[1]
		pwRuneRange := strings.Split(split[0], "-")
		pwRuneMin, _ := strconv.Atoi(pwRuneRange[0])
		pwRuneMax, _ := strconv.Atoi(pwRuneRange[1])
		pw := split[2]

		count := strings.Count(pw, pwRune)
		if pwRuneMin <= count && count <= pwRuneMax {
			valid++
		}
		continue
	}

	fmt.Printf("There are %d valid passwords\n", valid) // 414
}

/* 
Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy enforcement.

Given the same example list from above:

    1-3 a: abcde is valid: position 1 contains a and position 3 does not.
    1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
    2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.

How many passwords are valid according to the new interpretation of the policies?
*/
func (s *server) day2p2() {
	xs := s.parseFileToXStr("day2")

	var valid int

	for _, str := range xs {
		split := strings.Split(str, " ")
    split[1] = strings.TrimSuffix(split[1], ":")

    pw := split[2]
    
    pwRune := split[1] // rune required
    pwRuneByte := pwRune[0] // convert to byte for comparisons

    // Turn the rune positions into the bytes representing those runes so that I can compare them
		pwRunePos := strings.Split(split[0], "-")
		pwRunePos1, _ := strconv.Atoi(pwRunePos[0])
		pwRunePos2, _ := strconv.Atoi(pwRunePos[1])
    pwRunePos1 = pwRunePos1 - 1
		pwRunePos2 = pwRunePos2 - 1
    pwRuneByte1 := pw[pwRunePos1]
    pwRuneByte2 := pw[pwRunePos2]

    // Check pwRunePos1 & 2, check that they have 1 instance of pwRune
    condition := false
    if pwRuneByte1 == pwRuneByte || pwRuneByte2 == pwRuneByte {
      if pwRuneByte1 != pwRuneByte2 {
        condition = true
      }
    }
    
		if condition {
			valid++
		}
		continue
	}

	fmt.Printf("There are %d valid passwords\n", valid) // 413
}