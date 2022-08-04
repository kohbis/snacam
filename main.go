package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

const (
	// unexpectedSubcommand : subcommand is none of 'snake', 'camel' and 'pascal'
	unexpectedSubcommand = "expected 'snake', 'camel(lowercamel)' or 'pascal(uppercamel)' subcommand"
)

// CamelCaseType : Type of camelcase
type CamelCaseType int

const (
	// NONE : Not CamelCase
	NONE CamelCaseType = iota
	// UCC : Upper CamelCase
	UCC
	// LCC : Lower CamelCase
	LCC
)

func main() {
	var (
		snakeCmd  = flag.NewFlagSet("snake", flag.ExitOnError)
		camelCmd  = flag.NewFlagSet("camel", flag.ExitOnError)
		pascalCmd = flag.NewFlagSet("pascal", flag.ExitOnError)
	)

	if len(os.Args) < 2 {
		fmt.Println(unexpectedSubcommand)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "snake":
		snakeCmd.Parse(os.Args[2:])
		words := buildWords(snakeCmd.Args(), NONE)
		fmt.Println(strings.Join(words, "_"))
	case "camel", "lowercamel":
		camelCmd.Parse(os.Args[2:])
		words := buildWords(camelCmd.Args(), LCC)
		fmt.Println(strings.Join(words, ""))
	case "pascal", "uppercamel":
		pascalCmd.Parse(os.Args[2:])
		words := buildWords(pascalCmd.Args(), UCC)
		fmt.Println(strings.Join(words, ""))
	default:
		fmt.Println(unexpectedSubcommand)
		os.Exit(1)
	}
}

func buildWords(args []string, ccType CamelCaseType) []string {
	words := []string{}

	for _, arg := range args {
		for _, s := range regexp.MustCompile("[a-zA-Z][a-z]*").FindAllString(arg, -1) {
			word := strings.ToLower(s)
			words = append(words, word)
		}
	}

	if ccType == NONE {
		return words
	}
	return capitalizeWords(words, ccType)
}

func capitalizeWords(words []string, ccType CamelCaseType) []string {
	for i, word := range words {
		if i == 0 && ccType == LCC {
			continue
		}
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}

	return words
}
