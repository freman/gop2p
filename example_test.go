package gop2p_test

import (
	"fmt"
	"time"

	"github.com/lucasmenendez/gop2p"
)

func Example() {
	// Creating main node with debug mode equal to false. Then set individual
	// handlers.
	main := gop2p.InitNode(5001, false)
	// Wait for connections.
	defer main.Wait()

	// Set a connection handler
	main.On("connection", func(usr []byte) {
		fmt.Printf("[main handler] -> Connected: %s\n", string(usr))
	})

	// Set a message handler.
	main.On("message", func(msg []byte) {
		fmt.Printf("[main handler] -> Message: %s\n", string(msg))
	})

	// Set a disconnection handler
	main.On("disconnection", func(usr []byte) {
		fmt.Printf("[main handler] -> Disconnected: %s\n", string(usr))
	})

	// Creating peer on localhost 5002 port.
	go func() {
		// Wait for main node initialization.
		time.Sleep(time.Second)
		// Get main peer and create node in debug mode. To create an entry peer
		// manually, use CreatePeer function.
		entry := main.Self
		node := gop2p.InitNode(5002, true)
		defer node.Wait()

		// Connect to main node peer.
		node.Connect(entry)
		// Wait and broadcast message.
		time.Sleep(time.Second)
		node.Broadcast([]byte("Hello peers!"))
		// Wait and disconnect
		time.Sleep(2 * time.Second)
		node.Disconnect()
	}()

	// Create peer on localhost 5003 port.
	go func() {
		time.Sleep(time.Second)
		entry := main.Self

		node := gop2p.InitNode(5003, false)
		defer node.Wait()

		node.Connect(entry)
		time.Sleep(2 * time.Second)
		node.Disconnect()
	}()

	// Wait and broadcast. Broadcast fail is expected.
	time.Sleep(6 * time.Second)
	main.Broadcast([]byte("Hello peers!"))
	// Wait and disconnect.
	time.Sleep(2 * time.Second)
	main.Disconnect()

	// Output:[main handler] -> Connected: {"port":"5003","address":"192.168.0.37"}
	//[main handler] -> Connected: {"port":"5002","address":"192.168.0.37"}
	//[main handler] -> Message: Hello peers!
	//[main handler] -> Disconnected: {"port":"5003","address":"192.168.0.37"}
	//[main handler] -> Disconnected: {"port":"5002","address":"192.168.0.37"}
}
