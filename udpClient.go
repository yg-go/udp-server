package main

import (
	"context"
	"encoding/binary"
	"net"
	"runtime/debug"
	"sync"
)

type UDPCom struct {
	context context.Context
	address string
	port    string
}

type UDPEventHandler interface {
	OnMessage(payload []byte)
}

func NewUDPClient(ctx context.Context, address string, port string) *UDPCom {
	return &UDPCom{
		context: ctx,
		address: address,
		port:    port,
	}
}

func defaultPanicHandler(v interface{}) {
	log.Println(v)
	debug.PrintStack()
}

func UDPClientStart(com *UDPCom, handler UDPEventHandler, ifi *net.Interface) error {
	const MAX_BUF_SIZE int = 65535

	groupAddr := com.address + ":" + com.port
	addr, err := net.ResolveUDPAddr("udp", groupAddr)
	if err != nil {
		return err
	}
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	conn.SetReadBuffer(MAX_BUF_SIZE)
	defer conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if v := recover(); v != nil {
				defaultPanicHandler(v)
			}
		}()
		buf := make(Payload, MAX_BUF_SIZE)
		for {
			conn.ReadFromUDP(buf)
			size := int32(binary.LittleEndian.Uint16(buf[4:8])) + 8
			packet := buf[0:size]
			handler.OnMessage(packet)
		}
	}()
	wg.Wait()
	return nil
}
