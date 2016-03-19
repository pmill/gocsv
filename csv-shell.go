package main;

import (
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

func ExtractFirstNRowsShell(source string, destination string, maxRows int, randomise bool) (error) {
    defer un(trace("ExtractFirstNRowsGo"))
    
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
