package parser

import (
	"github.com/67in/crawler/c4/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all{
		result.Items = append(result.Items, string(c[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(c[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
