# This program calculates shortest substring containing strings from input array.
package main

import (
	"fmt"
	"strings"
	"sort"
)


const input = "this is a test.this is a programming test. a programming test this is"

type ValArray []int

func (a ValArray) Len() int           { return len(a) }
func (a ValArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ValArray) Less(i, j int) bool { return a[i] < a[j] }

func shortestString (input string, outputList []string) {
  strMatch := make(map[string]int)
  inputArray := strings.Split(input, " ")
  for i, val := range inputArray{
    for _, out := range outputList {
      if (val == out) {
        strMatch[val] = i
      }
    }
  }
  
  fmt.Println(strMatch)

  valArray := make (ValArray, 0)
  for _, val := range strMatch {
     valArray = append (valArray, val)
     fmt.Println(valArray)

  }
  sort.Sort(valArray)
  fmt.Println(valArray[0])


  for i:=valArray[0]; i<len(inputArray); i++ {
  fmt.Println(inputArray[i])
  fmt.Println(i)


  }

}


func main() {
outputList := []string{"this", "test", "is", "programming"}

shortestString(input, outputList)

}





