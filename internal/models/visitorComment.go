package models

import (
	"log"
	"strings"
)

type Visitor interface {
	VisitComment(*Comment)
}

type BadWordDetector struct {
	badWords []string
}

func (v *BadWordDetector) VisitComment(c *Comment) {
	for _, word := range v.badWords {
		if strings.Contains(strings.ToLower(c.Content), word) {
			log.Printf("Comment contains sensitive content. Not adding to blockchain.")
			c.Content = "" // Clear the content or handle it as you see fit
			return
		}
	}
}
