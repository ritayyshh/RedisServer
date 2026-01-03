package server

import (
	"net"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	// Start server in a goroutine
	go StartServer()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Send some data
	testData := []byte("Hello Server")
	n, err := conn.Write(testData)
	if err != nil {
		t.Fatalf("Failed to write data: %v", err)
	}

	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(testData), n)
	}

	t.Logf("Successfully sent %d bytes to server", n)
}

func TestMultipleConnections(t *testing.T) {
	// Start server in a goroutine (if not already running)
	go StartServer()
	time.Sleep(100 * time.Millisecond)

	// Test multiple concurrent connections
	for i := 0; i < 5; i++ {
		go func(id int) {
			conn, err := net.Dial("tcp", "localhost:6379")
			if err != nil {
				t.Errorf("Connection %d failed: %v", id, err)
				return
			}
			defer conn.Close()

			conn.Write([]byte("Test message"))
		}(i)
	}

	time.Sleep(500 * time.Millisecond)
}
