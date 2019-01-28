package parser

import (
	"awesomeProject/m/crawler/engine"
	"regexp"
)

const cityRe = ``

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "User" + string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
