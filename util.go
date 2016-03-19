package main;

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
    "time"
)

func trace(s string) (string, time.Time) {
    return s, time.Now()
}

func un(s string, startTime time.Time) {
    endTime := time.Now()
    fmt.Println("Elapsed time in seconds:", endTime.Sub(startTime))
}

func writeLines(rows [][]string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := csv.NewWriter(file)
  defer w.Flush()
  
  for _, row := range rows {
    w.Write(row)
  }
  
  return nil
}

func checkFileExists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func assertValidFilename(filename string) error {
    if len(strings.TrimSpace(filename)) == 0 {
        return fmt.Errorf("No filename submitted\r\n")
    }
    
    if !checkFileExists(filename) {
        return fmt.Errorf("%s does not exist\r\n", filename)
    }
    
    return nil
}
