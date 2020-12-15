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

func (s *server) day3p2() {
  xs := s.parseFileToXStr("day3")

  treeProduct := s.treesEncountered(xs, 1, 1)
  treeProduct *= s.treesEncountered(xs, 3, 1) // 280
  treeProduct *= s.treesEncountered(xs, 5, 1)
  treeProduct *= s.treesEncountered(xs, 7, 1)
  treeProduct *= s.treesEncountered(xs, 1, 2)

  fmt.Printf("\ntree product: %d\n", treeProduct)
}

// Given a slope of units right and units down, return an int that indicates '#' characters encountered
func (s *server) treesEncountered(treemap []string, right, down int) int {
  treesEncountered := 0
  column := 0
  for i := 0; i < len(treemap); i = i + down {
    s := treemap[i]

    // Check for a # rune at column of s
    rune := rune(s[column])
    if rune == 35 { // 35 is '#'
      treesEncountered++
    }
    column = column + right
    if column > 30 {
      column = column - 31
    }
  }

  fmt.Printf("right %d, down %d, trees: %d\n", right, down, treesEncountered)

  return treesEncountered
}