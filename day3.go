package main

import "fmt"

/*
.#..........#...#...#..#.......
.##X...#.#.##..###..#...#...#..
#.....X................#...#.#.
#.....#..X##.............#....#
*/
// In order to calculate the number of trees encountered following a given path in a repeating map, I will script a pathing algorithm that wraps to the beginning of the next row when needed.  I will then add to a count whether or not I have encountered a tree (#) or not.
func (s *server) day3p1() {
  xs := s.parseFileToXStr("day3")
  
  treesEncountered := 0
  column := 0
  for _, s := range xs {
    // Check for a # rune at column of s
    rune := rune(s[column])
    if rune == 35 { // 35 is '#'
      treesEncountered++
    }
    column = column + 3
    if column > 30 {
      column = column - 31
    }
  }
  
  fmt.Printf("trees encountered: %d\n", treesEncountered) // 280
}
