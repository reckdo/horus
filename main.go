package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	limit := 444
	for i := 1; i < limit; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Porta %d aberta\n", i)
	}
}
