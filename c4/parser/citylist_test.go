package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil{
		panic(err)
	}

	parseResult := ParseCityList(contents)

	const resultSize  = 470
	if len(parseResult.Requests) != resultSize{
		t.Errorf("result should have %d requests, but had %d",resultSize, len(parseResult.Requests))
	}
	if len(parseResult.Items) != resultSize{
		t.Errorf("result should have %d requests, but had %d",resultSize, len(parseResult.Requests))
	}
}
