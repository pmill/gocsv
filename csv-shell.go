package main;

import (
    "bytes"
    "encoding/csv"
    "fmt"
    "os/exec"
    "strconv"
    "strings"
)

func CountCSVRowsShell(source string) (int, error) {
    defer un(trace("CountCSVRowsShell"))
    
    err := assertValidFilename(source)
    if err != nil {
        return 0, err
    }
            
    cmd := exec.Command("wc", "-l", "<", source)
    cmdOut, err := cmd.Output()
    
    rowCountStrings := strings.Split(strings.TrimSpace(string(cmdOut)), " ")
    rowCountString := rowCountStrings[0]
    
    rowCount, _ := strconv.Atoi(rowCountString)
    
    return rowCount, nil
}

func ExtractFirstNRowsShell(source string, destination string, maxRows int) (error) {
    defer un(trace("ExtractFirstNRowsShell"))
    
    err := assertValidFilename(source)
    if err != nil {
        return err
    }
    
    if maxRows < 1 {
        maxRows = 1
    }
    
    fmt.Printf("Fetching %d rows from %s\r\n", maxRows, source)
    
    cmd := exec.Command("sed", "-n", "'1, 5000 p'", source, ">", destination)
    _ = cmd.Run()
    
    return nil
}

func ExtractRandomNRowsShell(source string, destination string, maxRows int) (error) {
    defer un(trace("ExtractRandomNRowsShell"))
    
    err := assertValidFilename(source)
    if err != nil {
        return err
    }
    
    if maxRows < 1 {
        maxRows = 1
    }
    
    fmt.Printf("Fetching %d random rows from %s\r\n", maxRows, source)
    
    cmd := exec.Command("shuf", "-n", fmt.Sprintf("%d", maxRows), source, ">", destination)
    _ = cmd.Run()
    
    return nil
}

func ExtractRowShell(source string, rowNumber int) ([]string, error) {
    defer un(trace("ExtractRowShell"))
    
    err := assertValidFilename(source)
    if err != nil {
        return nil, err
    }
    
    cmd := exec.Command("sed", "-n", fmt.Sprintf("%dp", rowNumber), source)
    cmdOut, _ := cmd.Output()
    
    b := bytes.NewBufferString(string(cmdOut))
    r := csv.NewReader(b)
    
    rows, err := r.ReadAll()
    if err != nil {
        return nil, err
    }
    
    if len(rows) == 0 {
        return nil, fmt.Errorf("No header found")
    }
    
    return rows[0], nil
}
