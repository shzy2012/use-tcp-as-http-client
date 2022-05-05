package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	// 使用tcp连接http服务端
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write(body())
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println(string(reply))
}

func body() []byte {

	boundary := randomBoundary()

	// body bytes
	var body bytes.Buffer
	// add file
	imgBytes, _ := ioutil.ReadFile("./pic.jpg")
	fmt.Fprintf(&body, "\r\n--%s\r\n", boundary) //开头--boundary
	fmt.Fprintf(&body, "%s\r\n", "Content-Disposition: form-data; name=\"files\"; filename=\"pic.jpg\"")
	fmt.Fprintf(&body, "%s\r\n\r\n", "Content-Type: image/jpg")
	body.Write(imgBytes)
	fmt.Fprintf(&body, "\r\n--%s--\r\n", boundary) //结束--boundary--

	var b bytes.Buffer
	fmt.Fprintf(&b, "%s\r\n", "POST /form?ID=0003 HTTP/1.1")
	fmt.Fprintf(&b, "%s\r\n", "Accept: */*")
	fmt.Fprintf(&b, "%s\r\n", "Host: localhost:8000")
	fmt.Fprintf(&b, "%s\r\n", "Accept-Encoding: gzip, deflate, br")
	fmt.Fprintf(&b, "%s\r\n", fmt.Sprintf("Content-Type: multipart/form-data; boundary=%s", boundary))
	fmt.Fprintf(&b, "%s\r\n", fmt.Sprintf("Content-Length: %v", len(body.Bytes())))
	fmt.Fprintf(&b, "%s\r\n", "")
	b.Write(body.Bytes())

	// fmt.Printf("%s\n", b.Bytes())
	return b.Bytes()
}

func randomBoundary() string {
	var buf [30]byte
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf[:])
}
