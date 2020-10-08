package udp

import (
	"net"
	"log"
	"strconv"
)

var receiver_ports = []string {"127.0.0.1:3002", "127.0.0.1:3003"}
var receiver_conns = [2]net.Conn {}

func Listen (port int) {
	addr, err := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(port))
	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatal(err)
	}

	// Close the connection
	// when Listen() returns
	defer conn.Close()

	// Create connections
	// for all recipients
	for i, str := range receiver_ports {
		dst, err := net.ResolveUDPAddr("udp", str)
		c, _ := net.DialUDP("udp", nil, dst)

		if err == nil {
			receiver_conns[i] = c
			log.Println(receiver_conns)
		} else {
			log.Println(err)
		}
	}

	for {
		buf := make([]byte, 4096)
		n, _, err := conn.ReadFromUDP(buf)

		if err != nil {
			log.Fatal(err)
			continue
		}

		// Create a slice containing
		// the actual packet only by
		// using the length returned
		// by ReadFromUDP
		conc := buf[0:n]

		// Send the packet to all
		// the recipients
		for _, conn := range receiver_conns {
			conn.Write(conc)
		}
	}
}