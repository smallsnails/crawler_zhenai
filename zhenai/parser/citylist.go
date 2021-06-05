package parser

import (
	"crawler_zhenai/defs"
	"crawler_zhenai/model"
	"regexp"
)

const cityListRule = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

func ParseCitylist(content []byte) defs.ParseResult {
	re := regexp.MustCompile(cityListRule)
	matches := re.FindAllStringSubmatch(string(content), -1)

	result := defs.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, defs.Request{
			//Url:       string(m[1]) + "," + string(m[2]),
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
		citylist := model.Citylist{
			Url:  string(m[1]),
			Name: string(m[2]),
		}
		//result.Items = append(result.Items, string(m[2]))
		result.Items = append(result.Items, citylist)
		result.Flag = "citylist"
	}
	return result
}
