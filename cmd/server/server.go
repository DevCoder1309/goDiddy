package server

import (
	"fmt"
	"godiddy/internal/adapters/decoder"
	"godiddy/internal/adapters/encode"
	"godiddy/internal/adapters/utils"
	"godiddy/internal/core/domain"
	"net"
	"os"
)

type packet struct {
	data []byte
	addr net.Addr
}

func Run() {
	ln, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ln.Close()

	c := make(chan packet)

	for {
		go listenAddr(ln, c)

		p := <-c
		
		// header, _, err := decoder.Decode(decoder.DecodeStruct{}, p.data)
		header, err := decoder.DecodeStruct.DecodeHeader(decoder.DecodeStruct{}, p.data)

		if err != nil {
			os.Exit(1)
		}

		var response *domain.Response

		questions, err:= decoder.DecodeStruct.DecodeQuestion(decoder.DecodeStruct{}, p.data, header.QDCount)
		response, err = utils.Pass(header, questions)
		if err != nil {
			os.Exit(1)
		}
		sentResponse := encode.EncodeResponse(response)
		go sendResponse(ln, p.addr, sentResponse)

		// if err != nil {
		// 	os.Exit(1)
		// }

		// sentResponse := encode.EncodeResponse(response)

		// go sendResponse(ln, p.addr, sentResponse)

		// fmt.Println(sentResponse)
	}
}

func listenAddr(ln net.PacketConn, c chan packet) {
	bs := make([]byte, 9999)

	n, addr, err := ln.ReadFrom(bs)
	if err != nil {
		fmt.Println(err)
		return
	}

	c <- packet{
		data: bs[:n],
		addr: addr,
	}
}

func sendResponse(conn net.PacketConn, addr net.Addr, sentResponse []byte) {
	_, err := conn.WriteTo(sentResponse, addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v\n", err)
	}
}