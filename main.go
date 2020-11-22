package main

import (
	"io"
	"log"
	"github.com/axelboberg/go-rtp-repeater/internal/tcp"
	"github.com/axelboberg/go-rtp-repeater/internal/udp"
	"github.com/axelboberg/go-rtp-repeater/internal/rmrp"
)

func onListener (listener io.Reader) {
	var cerr error
	for cerr == nil {
		str, err := tcp.ReadString(listener)
		cerr = err
		
		log.Print(str)

		parts, parseErr := rmrp.Parse(str)
		
		if parseErr != nil {
			log.Println(parseErr)
			cerr = parseErr
		} else {
			switch parts[0] {
				
			// ROUTE
			case "ROUTE":
				udp.Route(parts[1], parts[2])
				
			// ADD
			case "ADD":
				if parts[1] == "SERVER" {
					go udp.CreateServer(parts[2])
				} else {
					go udp.CreateClient(parts[2])
				}
			
			// REMOVE
			case "REMOVE":
				udp.Kill(parts[1])
				
			}
		}
	}
}

func main () {	
	tcp.CreateServer(":3000", onListener)
	log.Println("Go is alive")
}