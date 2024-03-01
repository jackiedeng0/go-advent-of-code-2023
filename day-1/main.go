// Day 1: Trebuchet?!
package main

import (
  "fmt"
  "log"
  "os"
  "bufio"
  "strings"
  "strconv"
  "errors"
)

func getCValue(s string) (int, error) {
  firstIndex := strings.IndexAny(s, "0123456789")
  if firstIndex == -1 {
    return 0, errors.New("No digits appear in line")
  }
  // lastIndex must exist if firstIndex does
  lastIndex := strings.LastIndexAny(s, "0123456789")
  cstr := string([]byte{s[firstIndex], s[lastIndex]})
  cValue, _ := strconv.Atoi(cstr)
  return cValue, nil
}

func main() {

  file, err := os.Open("in.txt")
  if err != nil { log.Fatal(err) }
  // Set new scanner with new lines splitter
  reader := bufio.NewReader(file)

  sum := 0
  for {
    line, err := reader.ReadString('\n')
    if err != nil { break }

    cValue, err := getCValue(line)
    if err != nil { fmt.Println(err); continue }
    sum += cValue
  }

  fmt.Printf("Sum of calibration values: %d\n", sum)

}

