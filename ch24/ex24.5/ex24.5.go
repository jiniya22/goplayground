package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func dining(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s가 밥을 먹으려합니다.\n", name)
		first.Lock()
		fmt.Printf("%s가 %s를 획득했습니다.\n", name, firstName)
		second.Lock()
		fmt.Printf("%s가 %s를 획득했습니다.\n", name, secondName)

		fmt.Printf("%s가 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go dining("A", fork, spoon, "fork", "spoon")
	go dining("B", spoon, fork, "spoon", "fork")

	wg.Wait()
}
