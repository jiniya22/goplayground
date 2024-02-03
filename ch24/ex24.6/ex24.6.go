package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type PowerJob struct {
	num int
}

func (pj *PowerJob) Do() {
	fmt.Printf("%d 거듭제곱 job start\n", pj.num)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d 거듭제곱 job end! 결과는 %d\n", pj.num, pj.num*pj.num)
}

func main() {
	var jobs [10]Job
	size := 10

	for i := 0; i < size; i++ {
		jobs[i] = &PowerJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()

}
