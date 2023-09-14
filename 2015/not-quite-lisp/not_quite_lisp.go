package main

import (
	"errors"
	"fmt"
	"os"
  "strings"
)

// Find floor by parsing direction with character ( to go up and ) to go down
func FindFloor(direction string) (int, error) {
  floor := 0
  for _, val := range direction {
    switch val {
    case '(':
      floor++
    case ')':
      floor--
    default:
      return 0, errors.New("Error, direction input is using other than ( and )")
    }
  }
  return floor, nil
}

func PosEnterToBasement(direction string) int {
  pos := 0
  floor := 0
  for idx, val := range direction {
    switch val {
    case '(':
      floor++
    case ')':
      floor--
    }
    if floor == -1 {
      pos = idx
      break
    }
  }
  return pos+1
}

func main() {
  file, err := os.ReadFile("./input.txt")
  if err != nil {
    println("Error when try reading file")
  }
  input := strings.Trim(string(file), " \n")
  floor, err := FindFloor(input)
  if err != nil {
    fmt.Print(err)
  } else {
    fmt.Printf("Floor at %d", floor)
  }

  print("\ninstruction to basement as position: ")
  println(PosEnterToBasement(input))
}
