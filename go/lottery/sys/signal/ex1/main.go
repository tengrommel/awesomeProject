package ex1

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func handleSignal(signal os.Signal) {
	fmt.Println("Got", signal)
}

func main() {
	var wg sync.WaitGroup
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		for {
			sig := <-sigs
			fmt.Println(sig)
			handleSignal(sig)
			wg.Done()
			break
		}
	}()
	wg.Wait()
}
