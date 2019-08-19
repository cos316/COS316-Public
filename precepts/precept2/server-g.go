/*****************************************************************************
 * Name: 
 * NetId:
 *
 * Description: A simple TCP server using the Go net API.   Publishes
 *              an IPV4 IP address.
 *
 * Usage:       ./server-g [server port]
 * 
 * Example:     ./server-g 8999 
 * 
 *****************************************************************************/

package main

import (
   "net"
   "os"
   "log"
   "bufio"
)

const RECV_BUFFER_SIZE = 2048

func main() {
	// get port number from command line
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./server-g [server port]")
	}
	server_port := os.Args[1]

	// Create a TCP socket and bind to socket’s IP address and port
	// Listen on socket for new connections
	var err    error
	var socket net.Listener 
	socket, err = net.Listen("tcp", "127.0.0.1:" + server_port)
	if err != nil {
		log.Fatal("Error in Listen")
	}
   writer := bufio.NewWriter(os.Stdout)
   for {
       // look for a client to connect
	   var connection net.Conn
	   connection, err = socket.Accept()
	   if err != nil {
		   socket.Close()
		   log.Fatal(err)
	   }
	   message := make([]byte, RECV_BUFFER_SIZE)
	   bytesReceived, _ := connection.Read(message)
	   writer.Write(message[:bytesReceived])
	   writer.Flush()
	   
	   connection.Close()
   }
}