/**
 * User: guoyao
 * Time: 06-28-2013 11:19
 * Blog: http://www.guoyao.me
 */

package servers

import (
	"fmt"
	"net"
	"encoding/gob"
//	"time"
	"io"
	"net/http"
	"chapter13/servers/rpc"
)

func Main() {
//	go server()
//	go func() {
//		for {
//			client()
//			time.Sleep(time.Second * 5)
//		}
//	}()

//	go server()
//	go client()

//	http.HandleFunc("/hello", hello)
//	http.ListenAndServe(":9000", nil)

//	http.Handle(
//		"/assets/",
//		http.StripPrefix(
//			"/assets/",
//			http.FileServer(http.Dir("assets")),
//		),
//	)

	rpc.Main()

//	var input string
//	fmt.Scanln(&input)
}

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// send the message
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

// http

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
		<html>
		<head>
		<title>Hello World</title>
		</head>
		<body>
		Hello World!
		</body>
		</html>`,
	)
}
