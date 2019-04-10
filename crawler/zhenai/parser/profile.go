package parser

import (
	"Golang-Basic/crawler/engine"
	"fmt"
	"regexp"
)

const personalInfo = `objectInfo`

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(personalInfo)
	matches := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	for _, match := range matches {
		info := string(match[1])
		fmt.Println(info)
	}
	return result
}
