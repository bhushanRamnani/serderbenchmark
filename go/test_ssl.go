package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func measureTLSHandshakeAndNetworkIOTime(host string, port string, iterations int) (float64, float64) {
	var totalTLSHandshakeTime float64
	var totalNetworkIOTime float64

	for i := 0; i < iterations; i++ {
		// Create a TCP connection to the server
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		defer conn.Close()

		// Create a TLS client and perform the TLS handshake
		config := &tls.Config{InsecureSkipVerify: true} // InsecureSkipVerify for testing purposes
		startTimeTLSHandshake := time.Now()
		tlsConn := tls.Client(conn, config)
		if err := tlsConn.Handshake(); err != nil {
			fmt.Printf("Error during TLS handshake: %v\n", err)
			continue
		}
		tlsHandshakeTime := time.Since(startTimeTLSHandshake).Seconds() * 1000 // Convert to milliseconds
		totalTLSHandshakeTime += tlsHandshakeTime

		// Simulate basic network I/O (sending and receiving data)
		startTimeNetworkIO := time.Now()
		tlsConn.Write([]byte("GET / HTTP/1.1\r\nHost: " + host + "\r\n\r\n"))
		response := make([]byte, 4096)
		_, err = tlsConn.Read(response)
		if err != nil {
			fmt.Printf("Error during network I/O: %v\n", err)
			continue
		}
		networkIOTime := time.Since(startTimeNetworkIO).Seconds() * 1000 // Convert to milliseconds
		totalNetworkIOTime += networkIOTime
	}

	return totalTLSHandshakeTime / float64(iterations), totalNetworkIOTime / float64(iterations)
}

func main() {
	// Specify the server information
	host := "benchmark-test-462.s3.us-west-2.amazonaws.com"
	port := "443"

	// Specify the number of iterations for testing
	iterations := 10

	// Measure TLS handshake and network I/O time
	averageTLSHandshakeTime, averageNetworkIOTime := measureTLSHandshakeAndNetworkIOTime(host, port, iterations)

	fmt.Printf("Avg TLS Handshake Time (%d iterations): %.6f milliseconds\n", iterations, averageTLSHandshakeTime)
	fmt.Printf("Avg Network I/O Time (%d iterations): %.6f milliseconds\n", iterations, averageNetworkIOTime)
}
