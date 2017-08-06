package main

import (
    "os"
    "log"
    "fmt"
    "strings"
)

func main() {
    args := os.Args[1:]
    total_lines := 0
    for _, filename := range args {
        file, err_open := os.Open(filename)
        if err_open != nil {
            log.Fatal(err_open)
        }

        buffer := 1024
        data := make([]byte, buffer)
        lines := 0
        for count, err := file.Read(data);
            count != 0;
            count, err = file.Read(data) {

            if err != nil {
                log.Fatal(err)
            }
            lines += strings.Count(string(data[:count]), "\n")
        }

        fmt.Printf("%8d %s\n", lines, filename)
        total_lines += lines
    }
    if len(args) > 1 {
            fmt.Printf("%8d %s\n", total_lines, "total")
    }
}
