package main

import (
  "os"
  "strings"
  "strconv"
)

func CalculateWrappingPaperDim(l, w, h int) int {
  lw := l*w
  wh := w*h
  hl := h*l
  
  extra := 0
  if lw <= wh && lw <= hl {
    extra = lw
  } else if wh <= lw && wh <= hl {
    extra = wh
  } else if hl <= lw && hl <= wh {
    extra = hl
  }
  return (2 * (lw+wh+hl)) + extra
}

func main() {
  file, err := os.ReadFile("./input.txt")
  if err != nil {
    println("Error when try reading file input")
    return
  }
  sum := 0
  inputs := strings.Split(string(file), "\n")
  for i := 0; i < len(inputs)-1; i++ {
    input := strings.Split(inputs[i], "x")
    l, _ := strconv.Atoi(input[0])
    w, _ := strconv.Atoi(input[1])
    h, _ := strconv.Atoi(input[2])
    sum += CalculateWrappingPaperDim(l, w, h)
  }
  println(sum)
}
