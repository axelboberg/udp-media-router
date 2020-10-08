package main

import (
	"fmt"
	"github.com/axelboberg/go-rtp-repeater/internal/udp"
	"github.com/axelboberg/go-rtp-repeater/internal/http"
)

func main () {
	udp.Listen(3001)
	http.Listen(3000)
	fmt.Println("Go is alive")
}