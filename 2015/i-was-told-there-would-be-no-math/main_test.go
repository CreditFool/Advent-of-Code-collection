package main

import ( 
  "testing"
  "github.com/stretchr/testify/assert"
)

type box struct {
  l int
  w int
  h int
}

var test_cases = []struct {
  name string
  input box 
  expected int
}{
  { 
    name: "2x3x4",
    input: box{2, 3, 4},
    expected: 58,
  },
  {
    name: "1x1x10",
    input: box{1, 1, 10},
    expected: 43,
  },
  { 
    name: "3x4x2",
    input: box{3, 4, 2},
    expected: 58,
  },
  { 
    name: "4x3x2",
    input: box{4, 3, 2},
    expected: 58,
  },
  { 
    name: "2x2x2",
    input: box{2, 2, 2},
    expected: 28,
  },
  {
    name: "10x10x1",
    input: box{10, 10, 1},
    expected: 250,
  },

}

func TestCalculateWrappingPaperDim(t *testing.T) {
  for _, test := range test_cases {
    t.Run(test.name, func(t *testing.T) {
      result := CalculateWrappingPaperDim(test.input.l, test.input.w, test.input.h)
      assert.Equal(t, test.expected, result)
    })
  }
}

func BenchmarkCalculateWrappingPaperDim(b *testing.B) {
  for i := 0; i < b.N; i++ {
    for _, test := range test_cases {
      CalculateWrappingPaperDim(test.input.l, test.input.w, test.input.h)
    }
  }
}

