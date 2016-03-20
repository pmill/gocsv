package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "io"
    "fmt"
)

// CountCSVRowsGo returns a count of the number of rows in the give csv file
func CountCSVRowsGo(source string) (int, error) {
    defer un(trace("CountCSVRowsGo"))
    
    err := assertValidFilename(source)
    if err != nil {
        return 0, err
    }
        
    f, _ := os.Open(source)
    r := csv.NewReader(bufio.NewReader(f))
    
    rowCount := 0
    
    for {
        _, err := r.Read()
        if err == io.EOF {
            break
        }
        
        rowCount++
    }
    
    return rowCount, nil
}

func ExtractFirstNRowsGo(source string, destination string, maxRows int, randomise bool) (error) {
    defer un(trace("ExtractFirstNRowsGo"))
    
    err := assertValidFilename(source)
    if err != nil {
        return err
    }
    
    if maxRows < 1 {
        maxRows = 1
    }
    
    fmt.Printf("Fetching %d rows from %s\r\n", maxRows, source)
    f, _ := os.Open(source)
    r := csv.NewReader(bufio.NewReader(f))
    
    rowCount := 0
    rows := [][]string{}
    
    for {
        if (rowCount >= maxRows) {
            break
        }
        
        row, err := r.Read()
        if err == io.EOF {
            break
        }
        
        rows = append(rows, row)
        rowCount++
    }
    
    err = writeLines(rows, destination)
    if err != nil {
        return err
    }
    
    return nil
}
