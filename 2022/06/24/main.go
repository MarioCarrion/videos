package main

import (
	"fmt"
	"sync"
)

func main() {
	c := Choreographer{}

	wgSubs := sync.WaitGroup{}
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wgSubs.Add(1)
		wg.Add(1)

		go func(val int64) {
			ch := c.Subscribe()

			wgSubs.Done()

			for msg := range ch {
				c.Add(val)
				fmt.Println(msg, val)
			}

			wg.Done()
		}(int64(i) + 1)
	}

	wgSubs.Wait()

	c.Publish("message")
	c.Close()

	wg.Wait()

	fmt.Println("Value", c.Value())
}

type Choreographer struct {
	totalMux sync.RWMutex
	total    int64

	chansMux sync.RWMutex
	chans    []chan string
}

// Subscribe returns a channel meant to used for receiving messages.
func (o *Choreographer) Subscribe() <-chan string {
	o.chansMux.Lock()
	defer o.chansMux.Unlock()

	ch := make(chan string, 1)

	o.chans = append(o.chans, ch)

	return ch
}

// Publish sends a message to the subscribed channels.
func (o *Choreographer) Publish(msg string) {
	o.chansMux.RLock()
	defer o.chansMux.RUnlock()

	for _, ch := range o.chans {
		ch <- msg
	}
}

// Close indicates all subscribers no more messages are going to be sent.
func (o *Choreographer) Close() {
	o.chansMux.Lock()
	defer o.chansMux.Unlock()

	for _, ch := range o.chans {
		close(ch)
	}

	o.chans = nil
}

//-

// Add updates the local value by adding the received one.
func (o *Choreographer) Add(value int64) {
	o.totalMux.Lock()
	defer o.totalMux.Unlock()

	o.total += value
}

// Value prints out the current value
func (o *Choreographer) Value() int64 {
	o.totalMux.RLock()
	defer o.totalMux.RUnlock()

	return o.total
}
