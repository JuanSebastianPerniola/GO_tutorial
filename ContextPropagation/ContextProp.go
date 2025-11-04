// Context Propagation
// set a time out of 2 seconds
// launch 3 workers each simulating a work of long running tasks with random  duration set duration 1-4 s
// worker -> listen to ctx.Done() to if is cancelled
// Print a mssg
// Wait for all go routine to finish before existin the main function

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// We set a variable context in args so can be propagated
// For better readiblity ill use context instead of ctx
func worker(context context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d: Starting work\n", id)
	workTime := time.Duration(rand.Intn(4)+1) * time.Second

	select {
	case <-time.After(workTime):
		fmt.Printf("Worker %d: Completed work in %v\n", id, workTime)
	case <-context.Done():
		fmt.Printf("Worker %d: cancelled %v\n", id, context.Err())
	}
}
func main() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Diferentes latencias de red o carga del sistema
	context, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var waitgroup sync.WaitGroup
	for i := 1; i <= 3; i++ { // launch 3 workers
		waitgroup.Add(1)
		go worker(context, i, &waitgroup)
	}
	waitgroup.Wait()
	fmt.Println("All workers finished")
}
