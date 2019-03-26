package main

import (
	"Golang-Basic/retriever/mock"
	"Golang-Basic/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever = mock.Retriever{"this is a fake baidu.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
