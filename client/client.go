package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"log"
	"time"
	"echo-server/thrift-example-with-go/example"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

//100000次调用总用时8.274608845
//QPS:12085.163404482355

//10000次调用总用时0.729874584
//QPS:13700.984003575057

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, _ := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := example.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()

	data := example.Data{Text: "hello,world aaaaa bbbbb!"}

	start := time.Now()
	var n = 100000
	for i := 1; i <= n; i++ {
		d, err := client.DoFormat(&data)
		fmt.Println(d.Text)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
	}
	end := time.Since(start).Seconds()
	fmt.Printf("%v次调用总用时%v\n", n, end)
	fmt.Printf("QPS:%v\n", float64(n)/end)
}
