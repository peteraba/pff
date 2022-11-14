package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const digits = "0123456789"

// multiSplit is similar to strings.Split, but it can split a string by multiple runes, traversing the string once only
func multiSplit(word string, delimiters map[rune]struct{}) []string {
	results := []string{}

	runes := []rune(word)
	from := 0
	for i := 0; i < len(runes); i++ {
		if _, ok := delimiters[runes[i]]; ok {
			if i > from {
				results = append(results, string(runes[from:i]))
			}
			from = i + 1
		}
	}

	return results
}

func matchAll(words, searchTerms []string, numsOkay bool) bool {
	for _, term := range searchTerms {
		found := false
		for _, word := range words {
			if word == term {
				found = true
				break
			}
			if numsOkay && strings.Trim(word, digits) == term {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func visit(searchTerms []string, delimiters []rune, numsOkay, caseSensitive bool) filepath.WalkFunc {
	dMap := make(map[rune]struct{}, len(delimiters))
	for _, sep := range delimiters {
		dMap[sep] = struct{}{}
	}

	if !caseSensitive {
		for i, st := range searchTerms {
			searchTerms[i] = strings.ToLower(st)
		}
	}

	return func(path string, f os.FileInfo, err error) error {
		fileName := f.Name()
		if !caseSensitive {
			fileName = strings.ToLower(fileName)
		}
		words := multiSplit(fileName, dMap)
		if matchAll(words, searchTerms, numsOkay) {
			fmt.Printf("%s\n", path)
		}

		return nil
	}
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("no current path found. err: %v", err)
	}

	root := flag.String("root", currentDir, "directory used to find files (current working directory by default)")
	delimiters := flag.String("delimiters", "|-. ", "characters used to find word boundaries")
	numsOkay := flag.Bool("numsOkay", true, "if true, words prefixed or postfixed with numbers will be found")
	caseSensitive := flag.Bool("caseSensitive", false, "if true, searches will be case sensitive")
	flag.Parse()

	searchTerms := flag.Args()
	if len(searchTerms) < 1 {
		log.Fatalf("no search term provided.")
	}

	delimiterRunes := []rune(*delimiters)

	err = filepath.Walk(*root, visit(searchTerms, delimiterRunes, *numsOkay, *caseSensitive))
	if err != nil {
		log.Fatalf("filepath.Walk() error. err: %v", err)
	}
}
