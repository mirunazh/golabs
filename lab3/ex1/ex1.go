// 1. Un client cu ceasul decalat cere de un server timpul corect printr-o rețea cu întârzieri simulate
// cu time.Sleep. Pe lângă θ, vom folosi și durata round-trip-ului, fără timpul de procesare al
// serverului:

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c *Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}

func measure(cClock, sClock *Clock) (theta, delta time.Duration) {
	// cerer
	t1 := cClock.Now()
	time.Sleep(3*time.Millisecond +
		time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := sClock.Now()
	// proces
	time.Sleep(5*time.Millisecond +
		time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := sClock.Now()
	// rasp
	time.Sleep(3*time.Millisecond +
		time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := cClock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	initialOffset := -5 * time.Second

	server := &Clock{offset: 0}
	client := &Clock{offset: initialOffset}

	// a
	theta, delta := measure(client, server)
	fmt.Println("theta ", theta)
	fmt.Println("delta", delta)
	fmt.Println("off init ", client.offset)

	client.offset += theta

	fmt.Println("off dupa corectie ", client.offset)

	// b
	client.offset = initialOffset
	minTheta, minDelta := measure(client, server)
	for i := 1; i < 10; i++ {
		client.offset = initialOffset
		th, de := measure(client, server)
		if th < minTheta {
			minTheta = th
		}

		if de < minDelta {
			minDelta = de
		}
	}

	fmt.Println("theta min: ", minTheta)
	fmt.Println("delta min: ", minDelta)

	// comparatie
	fmt.Println("dif theta: ", minTheta-theta)
	fmt.Println("dif delta: ", minDelta-delta)
}
