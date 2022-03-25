// package main

// import (
// 	"fmt"
// 	"net"
// 	"time"
// )

// func main() {
// 	limit := 444
// 	for i := 1; i < limit; i++ {
// 		address := fmt.Sprintf("scanme.nmap.org:%d", i)
// 		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
// 		if err != nil {
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("Porta %d aberta\n", i)
// 	}
// }

package main

import (
	"fmt"
	"net"
	"time"
)

func generateGoroutines(host string, ports, resultScanner chan int) {

	for p := range ports {

		address := fmt.Sprintf("%v:%d", host, p)

		conn, err := net.DialTimeout("tcp", address, 1*time.Second)

		if err != nil {

			resultScanner <- 0
			continue
		}

		conn.Close()

		resultScanner <- p
	}
}

func SenderPortsProcess(limitPort int, ports chan int) {

	for i := 1; i <= limitPort; i++ {

		ports <- i
	}
}

func getResultPortsOpen(limitPort int, results chan int) []int {
	var listOpenPorts []int

	for i := 0; i < limitPort; i++ {

		port := <-results

		if port != 0 {

			listOpenPorts = append(listOpenPorts, port)
		}
	}

	return listOpenPorts
}

func main() {

	var host string = "scanme.nmap.org"

	var limitPort int = 1024

	ports := make(chan int, limitPort)

	resultsScan := make(chan int)

	for i := 0; i < cap(ports); i++ {

		go generateGoroutines(host, ports, resultsScan)
	}

	go SenderPortsProcess(limitPort, ports)

	portsOpens := getResultPortsOpen(limitPort, resultsScan)

	close(ports)

	close(resultsScan)

	fmt.Printf("Endereço verificado - %v\n", host)

	for _, port := range portsOpens {
		fmt.Printf("%d aberta\n", port)
	}
}



