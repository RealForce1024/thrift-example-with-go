package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"ThriftDemo/example"
	"log"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
)


//10000次调用总用时1.039004364
//QPS:10000

//100000次调用总用时10.687179334
//QPS:10000



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
	var n = 10000
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
	fmt.Printf("QPS:%v\n", n/int(end))
}
