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
	// unexpectedSubcommand : subcommand is neighter 'snake' nor 'camel'
	unexpectedSubcommand = "expected 'snake' or 'camel' subcommand"
)

// CamelCaseType : Type of camelcase
type CamelCaseType int

const (
	// UCC : Upper CamelCase
	UCC CamelCaseType = iota
	// LCC : Lower CamelCase
	LCC
)

func main() {
	var (
		snakeCmd = flag.NewFlagSet("snake", flag.ExitOnError)
		camelCmd = flag.NewFlagSet("camel", flag.ExitOnError)
	)

	if len(os.Args) < 2 {
		fmt.Println(unexpectedSubcommand)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "snake":
		snakeCmd.Parse(os.Args[2:])
		words := buildWords(snakeCmd.Args(), false)
		fmt.Println(strings.Join(words, "_"))
	case "camel":
		camelCmd.Parse(os.Args[2:])
		words := buildWords(camelCmd.Args(), true)
		fmt.Println(strings.Join(words, ""))
	default:
		fmt.Println(unexpectedSubcommand)
		os.Exit(1)
	}
}

func buildWords(args []string, capitalize bool) []string {
	words := []string{}

	for _, arg := range args {
		for _, s := range regexp.MustCompile("[a-zA-Z][a-z]*").FindAllString(arg, -1) {
			word := strings.ToLower(s)
			words = append(words, word)
		}
	}

	if capitalize {
		return capitalizeWords(words, LCC)
	}
	return words
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
