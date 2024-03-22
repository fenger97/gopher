package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 6606})
	if err != nil {
		log.Printf("listen tcp error: %v", err)
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go handleTCPConn(conn)
	}
}

func handleTCPConn(conn *net.TCPConn) {
	defer func() {
		_ = conn.Close()
	}()
	log.Printf("handle tcp conn: local: %s remote: %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
	buff := [1024]byte{}

	for {
		n, err := conn.Read(buff[:])
		if err != nil {
			log.Printf("read error: %v", err)
			return
		}
		log.Printf("read data: [% x]", buff[:n])
	}
}
