package network

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"online/game"
)

const (
	Port = "8080"
)

// RunAsHost runs the game as server (host)
func RunAsHost(w, h int) (net.Conn, error) {
	// Get local IP of the host
	hostIP, err := getLocalIP()
	if err != nil {
		hostIP = "localhost"
	}

	// Return IP to be displayed on screen
	return nil, fmt.Errorf("waiting_for_connection:%s", hostIP)
}

// AcceptConnection accepts an incoming connection
func AcceptConnection() (net.Conn, error) {
	listener, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		return nil, fmt.Errorf("listen tcp :%s: %w", Port, err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		return nil, fmt.Errorf("accept connection: %w", err)
	}
	return conn, nil
}

// getLocalIP gets the local IP of the host
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no local IP found")
}

// RunAsClient runs the game as client
func RunAsClient(w, h int) (net.Conn, error) {
	fmt.Print("Enter host IP (default: localhost): ")
	var hostIP string
	fmt.Scanln(&hostIP)
	if hostIP == "" {
		hostIP = "localhost"
	}

	conn, err := net.Dial("tcp", hostIP+":"+Port)
	if err != nil {
		return nil, fmt.Errorf("dial tcp %s:%s: %w", hostIP, Port, err)
	}

	return conn, nil
}

// SendGameState sends the game state through the connection
func SendGameState(conn net.Conn, gs *game.GameState) {
	data, err := json.Marshal(gs)
	if err != nil {
		return
	}
	conn.Write(data)
	conn.Write([]byte("\n"))
}

// ReadInputFromNetwork reads player input from the network
func ReadInputFromNetwork(conn net.Conn, inputChan chan string) {
	reader := bufio.NewReader(conn)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			inputChan <- "quit"
			return
		}
		inputChan <- strings.TrimSpace(input)
	}
}

// ReadGameStateFromNetwork reads the game state from the network
func ReadGameStateFromNetwork(conn net.Conn, gs *game.GameState) error {
	reader := bufio.NewReader(conn)
	jsonState, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	var receivedState game.GameState
	if err := json.Unmarshal([]byte(jsonState), &receivedState); err != nil {
		return err
	}

	gs.Players = receivedState.Players
	gs.Bullets = receivedState.Bullets
	gs.IsGameOver = receivedState.IsGameOver
	gs.Winner = receivedState.Winner

	return nil
}
