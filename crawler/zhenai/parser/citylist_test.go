package parser

import (
	"Golang-Basic/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",contents)
	result := ParseCityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests,but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests,but had %d", resultSize, len(result.Items))
	}

}
