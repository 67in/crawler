package main

import (
	"github.com/67in/crawler/c4/engine"
	"github.com/67in/crawler/c4/parser"
)

func main()  {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
