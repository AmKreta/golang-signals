package main

import (
	"signal/effect"
	"signal/signal"
	"time"
)

func main() {
	val, setVal := signal.CreateSignal[int16](40)

	effect.CreateEffect(func() {
		println("Value is:", val())
	})

	time.Sleep(1 * time.Second)
	setVal(50)
	time.Sleep(1 * time.Second)
	setVal(60)
	time.Sleep(1 * time.Second)
	setVal(70)
}
