package adet

import (
	"strings"
	"unicode"
)

type Detector struct {
	Streets []Street
	Metros  []Metro
}

func (d *Detector) DetectAddress(text string) (DetectionResult, error) {
	text = fixPunctuation(text)
	streets := extractStreet(d.Streets, text)
	metros := extractMetros(d.Metros, text)
	return DetectionResult{
		Streets: streets,
		Metros:  metros,
	}, nil
}

type DetectionResult struct {
	Streets        []Street
	Metros         []Metro
	BuildingNumber string
}

type Street struct {
	Name string
	Type string
}

type Metro struct {
	Name           string
	DistanceMetres int
}

// fixPunctuation adds space after all punctuation characters (":", ".", ",")
func fixPunctuation(text string) string {
	var result string
	var symbols = []rune{':', '.', ',', '!'}
	for i, character := range text {
		result += string(character)
		if i+1 == len(text) {
			break
		}
		for _, s := range symbols {
			if character == s && !unicode.IsSpace(rune(text[i]+1)) {
				result += " "
			}
		}
	}
	return result
}

func findStreetByName(allStreets []Street, name string) Street {
	for _, s := range allStreets {
		if strings.Contains(s.Name, name) {
			return s
		}
	}
	return Street{}
}

func extractStreet(allStreets []Street, text string) []Street {
	words := strings.Fields(text)
	var res []Street
	for _, word := range words {
		if s := findStreetByName(allStreets, word); s.Name != "" {
			res = append(res, s)
		}
	}
	return res
}

func extractMetros(allMetros []Metro, text string) (res []Metro) {
	for _, metro := range allMetros {
		if strings.Contains(text, metro.Name) {
			res = append(res, metro)
		}
	}
	return
}
