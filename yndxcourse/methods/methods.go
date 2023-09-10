package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	timeStart     time.Time
	timeFixations []time.Duration
}

func (sw *Stopwatch) Start() {
	sw.timeStart = time.Now()
	sw.timeFixations = sw.timeFixations[:0]
}

func (sw *Stopwatch) SaveSplit() {
	sw.timeFixations = append(sw.timeFixations, time.Now().Sub(sw.timeStart))
}

func (sw Stopwatch) GetResults() []time.Duration {
	return sw.timeFixations
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
