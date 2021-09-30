package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Process ID", os.Getpid())

	listenForWork()

	<-waitToExit()

	fmt.Println("exiting")
}

func listenForWork() {
	const workersN int = 5

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM)

	//-

	workersC := make(chan struct{}, workersN)

	// 1) Listen for messages to process
	go func() {
		for {
			<-sc

			workersC <- struct{}{} // 2) Send to processing channel
		}
	}()

	go func() {
		var workers int

		for range workersC { // 3) Wait for messages to process
			workerID := (workers % workersN) + 1
			workers++

			fmt.Printf("%d<-\n", workerID)

			go func() { // 4) Process messages
				doWork(workerID)
			}()
		}
	}()
}

func waitToExit() <-chan struct{} {
	runC := make(chan struct{}, 1)

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, os.Interrupt)

	go func() {
		defer close(runC)

		<-sc
	}()

	return runC
}

func doWork(id int) {
	fmt.Printf("<-%d starting\n", id)

	time.Sleep(3 * time.Second)

	fmt.Printf("<-%d completed\n", id)
}
