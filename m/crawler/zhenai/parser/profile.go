package parser

import (
	"awesomeProject/m/crawler/engine"
	"awesomeProject/m/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(``)
var marriageRe = regexp.MustCompile(``)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRe)
	return engine.ParseResult{Items: []interface{}{profile}}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
