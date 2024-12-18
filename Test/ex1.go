package main

import (
	"fmt"
)

func pcat(catchan <-chan struct{}, dogchan chan<- struct{}) {
	defer func() {
		fmt.Println("close dogchan")
		close(dogchan)
	}()

	i := 0
	for range catchan 	{
		i++
		if i > 10 {
			return
		}
		
		fmt.Printf("cat: %v\n", i)
		dogchan <- struct{}{}

	}
}

func pdog(dogchan <-chan struct{}, fishchan chan<- struct{}) {
	defer func() {
		fmt.Println("close fishchan")
		close(fishchan)
	}()

	i := 0
	for range dogchan {
		i++
		fmt.Printf("dog: %v\n", i)
		fishchan <- struct{}{}
	}
}

func pfish(fishchan <-chan struct{}, catchan chan<- struct{}, endchan chan<- struct{}) {
	defer func() {
		fmt.Println("close catchan")
		close(catchan)
		fmt.Println("close endchan")
		close(endchan)
	}()


	i := 0
	for range fishchan {
		i++
		fmt.Printf("fish: %v\n", i)
		catchan <- struct{}{}
	}
}

func main() {

	catchan := make(chan struct{})
	dogchan := make(chan struct{})
	fishchan := make(chan struct{})
	endchan := make(chan struct{})

	go pcat(catchan, dogchan)
	go pdog(dogchan, fishchan)
	go pfish(fishchan, catchan, endchan)

	catchan <- struct{}{}
	<-endchan


}