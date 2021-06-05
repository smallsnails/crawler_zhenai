package parser

import (
	"crawler_zhenai/defs"
	"crawler_zhenai/model"
	"fmt"
	"regexp"
)

const cityRule = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(content []byte) defs.ParseResult {
	re := regexp.MustCompile(cityRule)
	matches := re.FindAllStringSubmatch(string(content), -1)
	fmt.Println(matches)
	result := defs.ParseResult{}
	for _, m := range matches {
		//nickname := string(m[2])
		result.Requests = append(result.Requests, defs.Request{
			//Url:       string(m[1]) + "," + string(m[2]),
			Url: string(m[1]),
			//ParseFunc: func(bytes []byte) defs.ParseResult {
			//	return ParseProfile(bytes, nickname)
			//},
			ParseFunc: defs.NilParser,
		})
		city := model.City{
			Url:  string(m[1]),
			Name: string(m[2]),
		}
		//result.Items = append(result.Items, string(m[2]))
		result.Items = append(result.Items, city)
		result.Flag = "city"
	}
	return result
}
