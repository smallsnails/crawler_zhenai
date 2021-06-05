package parser

import (
	"crawler_zhenai/defs"
	"fmt"
	"regexp"
)

const profileRule = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseProfile(content []byte, nickname string) defs.ParseResult {
	re := regexp.MustCompile(profileRule)
	matches := re.FindAllStringSubmatch(string(content), -1)
	fmt.Println(matches)
	result := defs.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, defs.Request{
			Url:       string(m[1]) + "," + string(m[2]),
			ParseFunc: defs.NilParser,
		})
		result.Items = append(result.Items, string(m[2]))
	}
	return result
}
