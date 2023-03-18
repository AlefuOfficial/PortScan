package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	maxWorkers := 1000
	wg := &sync.WaitGroup{}
	wg.Add(maxWorkers)

	for port := 0; port <= 65535; port++ {
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("localhost:%d", p)
			conn, err := net.Dial("tcp", address)
			if err == nil {
				fmt.Printf("Port %d is open\n", p)
				conn.Close()
			} else {
			}
			return
		}(port)
		if (port+1)%maxWorkers == 0 {
			wg.Wait()
			wg.Add(maxWorkers)
		}
	}
	wg.Wait()
}
