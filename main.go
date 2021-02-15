package main

import (
	"flag"
	"strings"
	"log"
	"io/ioutil"
	"regexp"
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

	var symbols = map[string]bool {
		"(": true,
		")": true,
		"=": true,
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"%": true,
		">": true,
		"<": true,
		"[": true,
		"]": true,
	}
	identifierChars := regexp.MustCompile(`[_A-Za-z]`)
	numberChars := regexp.MustCompile(`[0-9]`)
	whiteSpace := regexp.MustCompile(`\s`)
	var tokens []Token
	// TODO add proper error handling - line number and maybe also character number
	tokenizer := Tokenizer{rs, symbols, identifierChars, numberChars, whiteSpace, tokens}
	tokenizerOutput := tokenizer.generate()

	var ast []AstItem
	parser := Parser{tokenizerOutput, 0, ast}
	astOutput := parser.generate()

	env := Env{parent: nil, vars: make(map[string]interface{})}
	print := func(args ...interface{}) interface{} {
		for _, arg := range args {
			fmt.Println(arg)
		}
		return nil
	}
	env.add("print", print)
	evalUnderEnv(astOutput, env)
}
