package main

import (
    "flag"
    "fmt"
)

func main() {
    filename := flag.String("filename", "", "enter a csv filename")
    destination := flag.String("output", "", "enter a csv filename")
    maxRows := flag.Int("maxRows", 1, "enter the number of rows you want to extract")
    action := flag.String("action", "rowCount", "enter an action [rowCount,extractFirstNRows,extractHeader,randomRows]")
    flag.Parse()
    
    switch *action {
        case "rowCount":
            rowCount, err := CountCSVRowsGo(*filename)
            if err != nil {
                fmt.Printf("%s\r\n", err.Error())
            } else {
                fmt.Printf("GO: There are %d rows in %s\r\n", rowCount, *filename)
            }
            
            rowCount, err = CountCSVRowsShell(*filename)
            if err != nil {
                fmt.Printf("%s\r\n", err.Error())
            } else {
                fmt.Printf("SHELL: There are %d rows in %s\r\n", rowCount, *filename)
            }
        case "extractFirstNRows":
            err := ExtractFirstNRowsGo(*filename, *destination, *maxRows)
            if err != nil {
                fmt.Printf("%s\r\n", err.Error())
            }
            
            err = ExtractFirstNRowsShell(*filename, *destination, *maxRows)
            if err != nil {
                fmt.Printf("%s\r\n", err.Error())
            }
        case "extractHeader":
            header, err := ExtractRowShell(*filename, 1)
            if err != nil {
                fmt.Printf("Err: %s\r\n", err.Error())
            } else {
                fmt.Printf("%s", header)
            }
        case "randomRows":
            err := ExtractRandomNRowsShell(*filename, *destination, *maxRows)
            if err != nil {
                fmt.Printf("%s\r\n", err.Error())
            }
    }
}
