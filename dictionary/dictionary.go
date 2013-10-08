package dictionary

import (
	"io/ioutil"
	"strings"
)

type dictionary struct {
	words []string
}

const (
	searchWord   int = 1
	searchPrefix int = 0
)

func New() *dictionary {
	d := new(dictionary)

	rawDictionary, err := ioutil.ReadFile("./dictionary.txt")
	if err != nil {
		panic("There was an error reading the dictionary file!")
	}

	d.words = strings.Split(string(rawDictionary), "\n")

	return d
}

func (d *dictionary) ContainsWord(word string) bool {
	return d.containsElement(word, searchWord)
}

func (d *dictionary) ContainsPrefix(prefix string) bool {
	return d.containsElement(prefix, searchPrefix)
}

func (d *dictionary) containsElement(str string, searchType int) bool {
	high := len(d.words)
	low := 0

	var mid, prefixHigh int
	var element string

	for low <= high {
		mid = (high + low) / 2

		element = d.words[mid]

		if searchType == searchPrefix {
			if len(str) > len(element) {
				prefixHigh = len(element)
			} else {
				prefixHigh = len(str)
			}
			element = element[0:prefixHigh]
		}

		if element == str {
			return true
		} else {
			if element < str {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return false
}
