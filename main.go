package main

import (
	"crawler_zhenai/defs"
	"crawler_zhenai/engine"
	"crawler_zhenai/fetcher"
	"crawler_zhenai/scheduler"
	"crawler_zhenai/zhenai/parser"
	"fmt"
)

func main() {
	c := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
	}
	c.Run(defs.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCitylist,
	})
}

func test() {
	url := "https://www.zhenai.com/zhenghun"
	content, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	res := parser.ParseCitylist(content)
	fmt.Println(res)
}
