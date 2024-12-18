package main

import (
	"fmt"
	"sync"
)

func pcat(wg *sync.WaitGroup, catchan <-chan struct{}, dogchan chan<- struct{}) {
	defer func() {
		fmt.Println("close dogchan")
		close(dogchan)
	}()

	i := 0
	for range catchan 	{
		i++
		fmt.Printf("cat: %v\n", i)
		dogchan <- struct{}{}
		wg.Done()
	}

	wg.Done()
}

func pdog(wg *sync.WaitGroup, dogchan <-chan struct{}, fishchan chan<- struct{}) {
	defer func() {
		fmt.Println("close fishchan")
		close(fishchan)
	}()

	i := 0
	for range dogchan {
		i++
		fmt.Printf("dog: %v\n", i)
		fishchan <- struct{}{}
		wg.Done()
	}

	wg.Done()
}

func pfish(wg *sync.WaitGroup, fishchan <-chan struct{}) {
	i := 0
	for range fishchan {
		i++
		fmt.Printf("fish: %v\n", i)
		wg.Done()
	}

	wg.Done()
}

func main() {

	catchan := make(chan struct{})
	dogchan := make(chan struct{})
	fishchan := make(chan struct{})

	var wg sync.WaitGroup
	go pcat(&wg, catchan, dogchan)
	go pdog(&wg, dogchan, fishchan)
	go pfish(&wg, fishchan)

	for i:=0; i<10; i++ {
		wg.Add(3)
		catchan <- struct{}{}
		wg.Wait()
	}

	wg.Add(3)
	close(catchan)
	wg.Wait()

}