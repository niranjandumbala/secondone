package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var text = &clock.Field{SecondText: "tick", MinuteText: "tock", HourText: "bong"}
	wg := &sync.WaitGroup{}

	// added wait group to wait till the execution of
	wg.Add(1)
	output := make(chan string)

	// create a after channel to end the execution
	endCh := time.After(time.Hour * time.Duration(3))

	// ticker for hour
	hourC := time.NewTicker(time.Hour * time.Duration(1))

	// ticker for minute
	minC := time.NewTicker(time.Minute * time.Duration(1))

	//  ticker for second
	secC := time.NewTicker(time.Second * time.Duration(1))
	clockConfig := clock.ClConfig{
		EndCh:endCh,
		Txt:text,
		Wg:wg,
		HourC:hourC,
		MinC:minC,
		SecC:secC,
		Ch:output,

	}
	// start the clock
	go clockConfig.Clock()

	// start the input timer
	startInputTime := time.After(time.Minute * time.Duration(10))

	// init new input for reading
	go text.InitInput("newinput",startInputTime)

	//print output
	go func() {
		for {
			fmt.Print(<-output)
		}
	}()
	wg.Wait()
	close(output)
}
