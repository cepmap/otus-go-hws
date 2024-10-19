package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// var regexpPattern = regexp.MustCompile(`[^А-я-]|- `)
var regexpPattern = regexp.MustCompile(`(?m)(-|!|\.|,|\s-|\(|\))*\s+`)

type words struct {
	word  string
	count int
}

func Top10(inputString string) []string {
	stringSlice := transformString(inputString, regexpPattern)
	wordsDict := makeStructSlice(stringSlice)
	wordsDict = sortSlice(wordsDict)
	return getTopTen(wordsDict)
}

func transformString(inputString string, regexpPatter *regexp.Regexp) []string {
	return strings.Fields(strings.ToLower(regexpPatter.ReplaceAllString(inputString, " ")))
}

func makeStructSlice(inputSlice []string) []words {
	wordsMap := make(map[string]int)
	for _, word := range inputSlice {
		wordsMap[word]++
	}
	wordsDict := make([]words, 0)
	for key, value := range wordsMap {
		wordsDict = append(wordsDict, words{key, value})
	}
	return wordsDict
}

func sortSlice(inSlice []words) []words {
	sort.Slice(inSlice, func(i, j int) bool {
		if inSlice[i].count != inSlice[j].count {
			return inSlice[i].count > inSlice[j].count
		}
		return inSlice[i].word < inSlice[j].word
	})
	return inSlice
}

func getTopTen(inSlice []words) []string {
	var returnSlice []string
	if len(inSlice) != 0 {
		for i := 0; i < 10 && i < len(inSlice); i++ {
			returnSlice = append(returnSlice, inSlice[i].word)
		}
	}
	return returnSlice
}
