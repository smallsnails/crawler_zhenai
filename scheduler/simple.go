package scheduler

import "crawler_zhenai/defs"

type SimpleScheduler struct {
	workerChan chan defs.Request
}

func (s *SimpleScheduler) Submit(r defs.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) Config(c chan defs.Request) {
	s.workerChan = c
}
