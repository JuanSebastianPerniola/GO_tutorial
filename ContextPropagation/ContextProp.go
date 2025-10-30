// ConteztPropagation
package main

import (
	"fmt"
	"time"
)

func main() {
	go thread(1)
	go thread(2)
	fmt.Scanln()
}

func thread(id int) {
	for i := 1; i < 10; i++ {
		fmt.Print("Thread", id, " : ", i)
		time.Sleep(100 * time.Millisecond)
	}
}
