package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	log *logrus.Logger
)

func main() {
	dispatcherCtx, cancel := context.WithCancel(context.Background())
	payloadChan := make(chan Payload)

	defer func() {
		close(payloadChan)
	}()
	// RF_GROUP = 239.0.0.253
	// RF_PORT = 10051
	go DispatcherHeader(payloadChan, dispatcherCtx)

	cli := NewUDPClient(context.Background(), os.Getenv("RF_GROUP"), os.Getenv("RF_PORT"))
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Start UDP Client On %s:%s", cli.address, cli.port)
		err := UDPClientStart(cli, &RFSUDPHandler{dispatcherChan: payloadChan}, nil)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	}()

	wg.Wait()
	cancel()
}
