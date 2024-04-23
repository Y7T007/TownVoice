package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type Visitor interface {
	VisitComment(*Comment)
}

type BadWordDetector struct {
	BadWords []string
}

func NewBadWordDetector() *BadWordDetector {
	// Read the bad words from the JSON file
	badWordsBytes, err := ioutil.ReadFile("./internal/config/badWords.json")
	if err != nil {
		log.Fatalf("Failed to read bad words file: %v", err)
	}

	// Unmarshal the JSON into a slice
	var badWords []string
	err = json.Unmarshal(badWordsBytes, &badWords)
	if err != nil {
		log.Fatalf("Failed to parse bad words file: %v", err)
	}

	return &BadWordDetector{
		BadWords: badWords,
	}
}

func (v *BadWordDetector) VisitComment(c *Comment) {
	for _, word := range v.BadWords {
		if strings.Contains(strings.ToLower(c.Content), word) {
			log.Printf("Comment contains sensitive content. Not adding to blockchain.")
			c.Content = "" // Clear the content or handle it as you see fit
			return
		}
	}
}
