package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

func request(url string) string {
	conn, err := net.Dial("tcp", url)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//result, err := readFully(conn)
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	return string(result)
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		result.Write(buf[0:n])
	}

	return result.Bytes(), nil
}

func getIP(host string) string {
	ipAddr, err := net.ResolveIPAddr("ip", host)
	checkError(err)
	return ipAddr.String()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func testRequest() {
	content := request("guoyao.me:80")
	fmt.Println(content)
}

func testGetIP() {
	ip := getIP("guoyao.me")
	fmt.Println(ip)
}
func main() {
	//testRequest()
	testGetIP()
}
