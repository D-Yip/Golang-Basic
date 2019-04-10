package parser

import (
	"Golang-Basic/crawler/engine"
	"fmt"
	"regexp"
)

const personalInfo = `<div class="m-btn purple"[^>]*>(([^<]+))</div>`

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(personalInfo)
	matches := re.FindSubmatch(contents)
	for _, match := range matches {
		info := string(match[1])
		fmt.Println(info)
	}
}
