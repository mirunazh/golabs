// 2. Trei goroutine își trimit mesaje între ele prin canale. Ele își trimit mesaje conform topologiei inel
// P1 → P2 → P3 → P1. Fiecare dintre aceste evenimente este înregistrat folosind ceasuri logice
// Lamport. Pornind de la șablonul de mai jos, implementați cele 3 metode ale ceasului și stocați în
// main evenimentele ordonate crescător. Puteți folosi sort.Slice cu funcția less. Între evenimente,
// contoarele ceasurilor sunt incrementate aleator.

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type LamportClock struct {
	counter int
}

func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() {
	l.Tick()
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.Tick()
}

type Message struct {
	sender, ts int
}

type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter == b.internal_counter {
		return a.process < b.process
	}
	return a.internal_counter < b.internal_counter
}

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)
	ch3 := make(chan Message)

	eventCh := make(chan Event)

	// procesul p1
	go func() {
		clock := LamportClock{}

		for i := 0; i < rand.Intn(5)+1; i++ {
			clock.Tick()
			eventCh <- Event{1, clock.counter, "internal"}
		}

		clock.Send()
		eventCh <- Event{1, clock.counter, "send"}
		ch1 <- Message{1, clock.counter}

		msg := <-ch3
		clock.Receive(msg.ts)
		eventCh <- Event{1, clock.counter, "recv"}
	}()

	// procesul p2
	go func() {
		clock := LamportClock{}

		for i := 0; i < rand.Intn(5)+1; i++ {
			clock.Tick()
			eventCh <- Event{2, clock.counter, "internal"}
		}

		msg := <-ch1
		clock.Receive(msg.ts)
		eventCh <- Event{2, clock.counter, "recv"}

		clock.Send()
		eventCh <- Event{2, clock.counter, "send"}
		ch2 <- Message{2, clock.counter}
	}()

	// procesul p3
	go func() {
		clock := LamportClock{}

		for i := 0; i < rand.Intn(5)+1; i++ {
			clock.Tick()
			eventCh <- Event{3, clock.counter, "internal"}
		}

		msg := <-ch2
		clock.Receive(msg.ts)
		eventCh <- Event{3, clock.counter, "recv"}

		clock.Send()
		eventCh <- Event{3, clock.counter, "send"}
		ch3 <- Message{3, clock.counter}
	}()

	events := []Event{}
	recvCount := 0

	// adunam ev pana avem 3
	for recvCount < 3 {
		e := <-eventCh
		events = append(events, e)
		if e.kind == "recv" {
			recvCount++
		}
	}

	// ord
	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})

	for _, e := range events {
		fmt.Println(e.process, e.internal_counter, e.kind)
	}
}
