package gop2p_test

import (
	"fmt"
	"time"

	"github.com/lucasmenendez/gop2p"
)

func Example() {
	// Creating main node in debug mode.
	main := gop2p.InitNode(5001, true)
	// Wait for connections.
	defer main.Wait()

	// Set a message handler.
	main.OnMessage(func(message []byte) {
		fmt.Printf("\t\t-> %s\n", string(message))
	})

	// Creating peer on localhost 5002 port.
	go func() {
		// Get main peer and create node in debug mode. To create an entry peer manually, use CreatePeer function.
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

		node := gop2p.InitNode(5003, true)
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

	// Output: 2018/08/10 23:13:37 [192.168.0.37:5001] Start event loop...
	// 2018/08/10 23:13:37 [192.168.0.37:5001] Listen at 192.168.0.37:5001
	// 2018/08/10 23:13:38 [192.168.0.37:5003] Listen at 192.168.0.37:5003
	// 2018/08/10 23:13:38 [192.168.0.37:5003] Start event loop...
	// 2018/08/10 23:13:38 [192.168.0.37:5003] Connecting to [192.168.0.37:5001]
	// 2018/08/10 23:13:38 [192.168.0.37:5002] Listen at 192.168.0.37:5002
	// 2018/08/10 23:13:38 [192.168.0.37:5002] Start event loop...
	// 2018/08/10 23:13:38 [192.168.0.37:5002] Connecting to [192.168.0.37:5001]
	// 2018/08/10 23:13:38 [192.168.0.37:5001] Connected to [192.168.0.37:5003]
	// 2018/08/10 23:13:38 [192.168.0.37:5003] Connected to [192.168.0.37:5001]
	// 2018/08/10 23:13:38 [192.168.0.37:5001] Connected to [192.168.0.37:5002]
	// 2018/08/10 23:13:38 [192.168.0.37:5002] Connected to [192.168.0.37:5003]
	// 2018/08/10 23:13:38 [192.168.0.37:5003] Connected to [192.168.0.37:5002]
	// 2018/08/10 23:13:38 [192.168.0.37:5002] Connected to [192.168.0.37:5001]
	// 2018/08/10 23:13:39 [192.168.0.37:5002] Message sended: 'Hola'
	// 2018/08/10 23:13:39 [192.168.0.37:5003] Message received: 'Hola'
	//                 -> Hola
	// 2018/08/10 23:13:39 [192.168.0.37:5001] Message received: 'Hola'
	// 2018/08/10 23:13:40 [192.168.0.37:5003] Disconnecting...
	// 2018/08/10 23:13:40 [192.168.0.37:5001] Disconnected From [192.168.0.37:5003]
	// 2018/08/10 23:13:40 [192.168.0.37:5002] Disconnected From [192.168.0.37:5003]
	// 2018/08/10 23:13:41 [192.168.0.37:5002] Disconnecting...
	// 2018/08/10 23:13:41 [192.168.0.37:5001] Disconnected From [192.168.0.37:5002]
	// 2018/08/10 23:13:43 [192.168.0.37:5001] Broadcasting aborted. Empty network!
	// 2018/08/10 23:13:45 [192.168.0.37:5001] Disconnecting...
}