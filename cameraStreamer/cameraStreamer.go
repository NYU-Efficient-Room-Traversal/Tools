package cameraStreamer // import github.com/NYU-Efficient-Room-Traversal/Tools/cameraStreamer

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"net"
	"os"
	"strings"
)

const (
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

var address string

func handleRequest(conn net.Conn, ch chan<- image.Image) {
	for {

		// Get base64 string from TCP Client
		str, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			fmt.Println("Closing Connection")
			return
		}

		// Remove whitespace and decode into jpeg image
		str = strings.TrimSpace(str)
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
		img, err := png.Decode(reader)
		if err != nil {
			fmt.Println("ERR:", err)
		} else {
			fmt.Println("Got Image")
			ch <- img
		}
	}

	fmt.Println("Closing Connection")
	return
}

func Open(ch chan<- image.Image) {
	// Attach Channel
	channel = ch

	// Get IP Address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				address = ipnet.IP.String()
			}
		}
	}

	// Begin listening
	listen, err := net.Listen(CONN_TYPE, address+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error establishing listener: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("TCP listener open on ", address, ":", CONN_PORT)
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("New Connection: ", conn)
		go handleRequest(conn, ch)
	}
}

func main() {
	c := make(chan image.Image, 10)
	go Open(c)
	for {
		i := <-c
		if i != nil {
			fmt.Println("get")
		}
	}
}
