package main

import (
	"Golang-Basic/retriever/mock"
	"Golang-Basic/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "abc",
			"ato":  "ye",
		})
}

func session(s RetrieverPoster) string {
	s.Post(url,
		map[string]string{
			"contents": "another faked baidu.com",
		})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake baidu.com"}
	r = &retriever
	r = real.Retriever{}
	fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}
