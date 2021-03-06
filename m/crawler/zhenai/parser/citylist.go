package parser

import (
	"awesomeProject/m/crawler/engine"
	"regexp"
)

const cityListRe = ``

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "City" + string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
