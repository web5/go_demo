package nets

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var serverHost = flag.String("serverHost", "localhost", "host")
var serverPort = flag.String("serverPort", "8999", "port")

func Server() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *serverHost+":"+*serverPort)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println("Error listening", err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		handleCleint(conn)
	}
}

func handleCleint(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of", err.Error())
		return
	}

	daytime := time.Now().Unix()
	fmt.Println(n, remoteAddr)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(daytime))
	conn.WriteToUDP(b, remoteAddr)
}
