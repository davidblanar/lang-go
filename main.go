package main

import (
	"flag"
	"strings"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
    filePath := flag.String("f", "", "Path to .l file to interpret")
    flag.Parse()
    if *filePath == "" {
		log.Fatal("Missing file path")
    }

    filePathArr := strings.Split(*filePath, ".")
    ext := filePathArr[len(filePathArr) - 1]
    if ext != "l" {
    	log.Fatal(fmt.Sprintf("Invalid file type '%s'", ext))	
    }

    fileContent, err := ioutil.ReadFile(*filePath)
    if err != nil {
        log.Fatal(err)
    }

    input := string(fileContent)
    rs := ReadStream{input, 0}
    fmt.Println(rs)
    // fmt.Println(TOKEN_IDENTIFIER)
}
