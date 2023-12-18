package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	host "github.com/libp2p/go-libp2p-host"
	libp2p "github.com/libp2p/go-libp2p"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Configure the host
	host, err := libp2p.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Obtain the host's multiaddress
	hostAddr := host.Addrs()[0]
	fmt.Printf("Host address: %s\n", hostAddr)

	// Start a background goroutine to handle incoming connections
	go handleIncomingConnections(ctx, host)

	// Wait for a termination signal
	waitForTerminationSignal()
}

func handleIncomingConnections(ctx context.Context, host host.Host) {
	for {
		select {
		case <-ctx.Done():
			return
		case conn := <-host.Network().Notify().NewPeers:
			fmt.Printf("New connection established: %s\n", conn)
		}
	}
}
