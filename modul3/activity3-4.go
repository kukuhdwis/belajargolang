package main

import (
	"fmt"
	"sync"
	"time"
)

// Chopstick represents a chopstick shared between adjacent philosophers
type Chopstick struct {
	sync.Mutex
}

// Host coordinates permission so no more than 2 philosophers eat concurrently
type Host struct {
	request  chan chan struct{}
	finished chan struct{}
	stop     chan struct{}
}

// NewHost creates and initializes a Host instance
func NewHost() *Host {
	return &Host{
		request:  make(chan chan struct{}),
		finished: make(chan struct{}),
		stop:     make(chan struct{}),
	}
}

// Start runs the host logic in its own goroutine
func (h *Host) Start() {
	go func() {
		var queue []chan struct{}
		eatingCount := 0

		for {
			// Grant permission to queued philosophers if capacity allows
			for len(queue) > 0 && eatingCount < 2 {
				reply := queue[0]
				queue = queue[1:]
				eatingCount++
				reply <- struct{}{}
			}

			select {
			case <-h.stop:
				return
			case reply := <-h.request:
				if eatingCount < 2 {
					eatingCount++
					reply <- struct{}{}
				} else {
					queue = append(queue, reply)
				}
			case <-h.finished:
				eatingCount--
			}
		}
	}()
}

// RequestPermission asks the host for permission to eat (blocks until granted)
func (h *Host) RequestPermission() {
	reply := make(chan struct{})
	h.request <- reply
	<-reply
}

// FinishedEating notifies the host that a philosopher has finished eating
func (h *Host) FinishedEating() {
	h.finished <- struct{}{}
}

// Stop terminates the host goroutine
func (h *Host) Stop() {
	close(h.stop)
}

// Philosopher represents one of the 5 dining philosophers
type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	host           *Host
}

// dine executes the eating routine 3 times for a philosopher
func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		// 1. Request permission from the host (max 2 concurrent eaters)
		p.host.RequestPermission()

		// 2. Pick up chopsticks in any order (left then right, not lowest-numbered first)
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		// 3. Print starting to eat (after obtaining locks)
		fmt.Printf("starting to eat %d\n", p.id)

		// Simulate eating time
		time.Sleep(10 * time.Millisecond)

		// 4. Print finishing eating (before releasing locks)
		fmt.Printf("finishing eating %d\n", p.id)

		// 5. Release chopsticks
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		// 6. Notify host that eating is complete
		p.host.FinishedEating()

		// Optional small pause between meals
		time.Sleep(5 * time.Millisecond)
	}
}

func main() {
	const numPhilosophers = 5

	// Create 5 chopsticks
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// Initialize and start the host in its own goroutine
	host := NewHost()
	host.Start()

	// Create 5 philosophers numbered 1 through 5
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
			host:           host,
		}
	}

	// Launch all 5 philosophers concurrently
	var wg sync.WaitGroup
	wg.Add(numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		go philosophers[i].dine(&wg)
	}

	// Wait for all 5 philosophers to complete eating 3 times each
	wg.Wait()

	// Stop host goroutine
	host.Stop()
}
