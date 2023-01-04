package main

import (
  "fmt"
  "path/filepath"
)

func main() {
  go files, _ := filepath.Glob("*.txt")
  fmt.Printf("%q\n", files)
}
}
