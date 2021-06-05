package engine

import (
	"crawler_zhenai/defs"
	"crawler_zhenai/fetcher"
	"crawler_zhenai/model"
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
}

type Scheduler interface {
	Submit(defs.Request)
	Config(chan defs.Request)
}

func (c *ConcurrentEngine) Run(seeds ...defs.Request) {
	in := make(chan defs.Request)
	out := make(chan defs.ParseResult)
	c.Scheduler.Config(in)

	for i := 0; i < 10; i++ {
		createWorker(in, out)
	}

	for _, seed := range seeds {
		//in <- seed
		//go func() {
		//	in <- seed
		//}()
		c.Scheduler.Submit(seed)
	}

	for {
		result := <-out
		for _, req := range result.Requests {
			//fmt.Println(i, req)
			//in <- req
			//go func() {
			//	in <- req
			//}()
			c.Scheduler.Submit(req)

			//name := runtime.FuncForPC(reflect.ValueOf(req.ParseFunc).Pointer()).Name()
			//if strings.Contains(name, "ParseCitylist") {
			//
			//} else if strings.Contains(name, "ParseCity") {
			//	model.InsertCityList(arr[0], arr[1])
			//	c.Scheduler.Submit(req)
			//} else if strings.Contains(name, "ParseProfile") {
			//	val := string(req.Url)
			//	arr := strings.Split(val, ",")
			//	fmt.Println(arr)
			//	model.InsertCity(arr[0], arr[1])
			//} else {
			//
			//}
		}
		flag := result.Flag
		for _, item := range result.Items {
			if flag == "citylist" {
				list := item.(model.Citylist)
				model.InsertCityList(list.Url, list.Name)
			} else if flag == "city" {
				list := item.(model.City)
				model.InsertCity(list.Url, list.Name)
			}
		}
	}

}

func createWorker(in chan defs.Request, out chan defs.ParseResult) {
	go func() {
		for {
			req := <-in
			result, err := worker(req)
			if err != nil {
				fmt.Println(err)
				continue
			}
			out <- result
		}
	}()
}

func worker(request defs.Request) (defs.ParseResult, error) {
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		fmt.Println(err)
		return defs.ParseResult{}, err
	}
	return request.ParseFunc(content), nil
}
