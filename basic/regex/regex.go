package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is yeshuijin@outlook.com.cn aaa
email1 is abc@sinotrans.com
email2 is abc@y2t.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
