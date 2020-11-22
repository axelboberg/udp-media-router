package udp

import (
	"net"
	"log"
)

var conns = Connections{conns: make(map[string]*net.UDPConn)}
var routing = Routing{up: make(map[string]string), down: make(map[string]map[string]bool)}

func dataLoop (addr string) {
	for {
		conn, ok := conns.Get(addr)
		if !ok {
			break
		}
		
		buf := make([]byte, 4096)
		n, _, err := conn.ReadFromUDP(buf)
		
		if err != nil {
			log.Print(err)
			continue
		}

		// Create a slice containing
		// only the actual packet by
		// using the length returned
		// by ReadFromUDP
		conc := buf[0:n]
		
		recvs := routing.Destinations(addr)
		
		for _, recv := range recvs {
			recv_conn, recv_ok := conns.Get(recv);
			if recv_ok {
				recv_conn.Write(conc)
			}	
		}
	}
}

func CreateServer (port string) {
	if _, ok := conns.Get(port); ok {
		log.Print("[UDP] Port", port, "is already in use")
		return
	}
	
	addr, err := net.ResolveUDPAddr("udp", port)
	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Print(err)
		return
	}

	conns.Add(port, conn)
	dataLoop(port)
}

func CreateClient (addr string) {
	dst, err := net.ResolveUDPAddr("udp", string(addr))
	conn, _ := net.DialUDP("udp", nil, dst)
	
	if err != nil {
		log.Println(err)
		return
	}
	
	conns.Add(addr, conn)
	dataLoop(addr)
}

func Kill (addr string) {
	if conn, ok := conns.Get(addr); ok {
		conns.Remove(addr)
		conn.Close()
	}
}

func Route (from string, to string) {
	routing.Route(from, to)
}