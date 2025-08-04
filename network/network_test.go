package network

import (
	"encoding/json"
	"net"
	"online/game"
	"strings"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	ip, err := getLocalIP()
	if err != nil {
		t.Fatalf("Failed to get local IP: %v", err)
	}

	if ip == "" {
		t.Error("Local IP should not be empty")
	}

	// Verify that it's a valid IP
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		t.Errorf("Invalid IP address: %s", ip)
	}

	t.Logf("Local IP found: %s", ip)
}

func TestRunAsHost(t *testing.T) {
	// test that RunAsHost returns the expected error format
	conn, err := RunAsHost(80, 24)

	if conn != nil {
		t.Error("RunAsHost should return nil connection")
	}

	if err == nil {
		t.Error("RunAsHost should return an error")
	}

	// check that error contains IP
	if !strings.Contains(err.Error(), "waiting_for_connection:") {
		t.Error("Error should contain 'waiting_for_connection:' prefix")
	}
}

func TestSendGameState(t *testing.T) {
	// TODO: implement proper connection mocking
	// skipping this test since it blocks on network operations
	t.Skip("Skipping test that requires network mocking")
}

func TestReadGameStateFromNetwork(t *testing.T) {
	// TODO: implement proper connection mocking
	// skipping this test since it blocks on network operations
	t.Skip("Skipping test that requires network mocking")
}

func TestReadInputFromNetwork(t *testing.T) {
	// TODO: implement proper connection mocking
	// skipping this test since it blocks on network operations
	t.Skip("Skipping test that requires network mocking")
}

func TestGameStateSerialization(t *testing.T) {
	gs := game.InitGame(true, 80, 24)

	// Simulate some state changes
	gs.Players[0].X = 10.5
	gs.Players[1].Y = 15.2

	// Create a test bullet
	bullet := &game.Bullet{
		X:       40,
		Y:       20,
		Speed:   1.0,
		OwnerID: 1,
	}
	gs.Bullets = append(gs.Bullets, bullet)

	// Serialize the state
	data, err := json.Marshal(gs)
	if err != nil {
		t.Fatalf("Failed to marshal GameState: %v", err)
	}

	// Deserialize the state
	var newGs game.GameState
	err = json.Unmarshal(data, &newGs)
	if err != nil {
		t.Fatalf("Failed to unmarshal GameState: %v", err)
	}

	// Verify that the data was preserved
	if len(newGs.Players) != len(gs.Players) {
		t.Errorf("Player count mismatch: expected %d, got %d", len(gs.Players), len(newGs.Players))
	}

	if len(newGs.Bullets) != len(gs.Bullets) {
		t.Errorf("Bullet count mismatch: expected %d, got %d", len(gs.Bullets), len(newGs.Bullets))
	}

	if newGs.Players[0].X != gs.Players[0].X {
		t.Errorf("Player 1 X position mismatch: expected %f, got %f", gs.Players[0].X, newGs.Players[0].X)
	}
}

func TestIPFunctionality(t *testing.T) {
	ip, err := getLocalIP()
	if err != nil {
		t.Fatalf("Failed to get local IP: %v", err)
	}

	if ip == "" {
		t.Error("Local IP should not be empty")
	}

	// Verify that it's a valid IP
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		t.Errorf("Invalid IP address: %s", ip)
	}

	t.Logf("Local IP found: %s", ip)
}
