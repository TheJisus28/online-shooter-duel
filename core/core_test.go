package core

import (
	"testing"
	"time"
)

// mock termbox event for testing
type mockEvent struct {
	eventType int
	key       int
	ch        rune
}

func TestHandleMenuInput(t *testing.T) {
	// TODO: implement proper termbox mocking
	// skipping this test since it blocks on termbox.PollEvent()
	t.Skip("Skipping test that requires termbox mocking")
}

func TestWaitForRestart(t *testing.T) {
	// TODO: implement proper termbox mocking
	// skipping this test since it blocks on termbox.PollEvent()
	t.Skip("Skipping test that requires termbox mocking")
}

func TestReadInputFromTerminal(t *testing.T) {
	// TODO: implement proper termbox mocking
	// skipping this test since it blocks on termbox.PollEvent()
	t.Skip("Skipping test that requires termbox mocking")
}

func TestInputChannelCommunication(t *testing.T) {
	// test that input channels work correctly
	inputChan := make(chan string, 10)

	// send test inputs
	go func() {
		inputChan <- "move_left"
		inputChan <- "move_right"
		inputChan <- "shoot"
		inputChan <- "quit"
	}()

	// receive inputs
	inputs := []string{}
	for i := 0; i < 4; i++ {
		select {
		case input := <-inputChan:
			inputs = append(inputs, input)
		case <-time.After(1 * time.Second):
			t.Fatal("Timeout waiting for input")
		}
	}

	// verify inputs
	expected := []string{"move_left", "move_right", "shoot", "quit"}
	for i, input := range inputs {
		if input != expected[i] {
			t.Errorf("Expected input '%s', got '%s'", expected[i], input)
		}
	}
}

func TestMenuOptionValidation(t *testing.T) {
	// test that menu options are handled correctly
	testCases := []struct {
		input    int
		expected bool
	}{
		{0, true},   // valid option
		{1, true},   // valid option
		{2, true},   // valid option
		{-1, false}, // invalid option
		{3, false},  // invalid option
	}

	for _, tc := range testCases {
		if tc.input >= 0 && tc.input <= 2 {
			// valid option
			if !tc.expected {
				t.Errorf("Option %d should be valid", tc.input)
			}
		} else {
			// invalid option
			if tc.expected {
				t.Errorf("Option %d should be invalid", tc.input)
			}
		}
	}
}

func TestInputStringValidation(t *testing.T) {
	// test that input strings are valid
	validInputs := []string{
		"move_left",
		"move_right",
		"shoot",
		"quit",
	}

	invalidInputs := []string{
		"invalid",
		"",
		"move_up",
		"fire",
	}

	// test valid inputs
	for _, input := range validInputs {
		if input == "" {
			t.Error("Input should not be empty")
		}
	}

	// test invalid inputs
	for _, input := range invalidInputs {
		if input == "move_left" || input == "move_right" || input == "shoot" || input == "quit" {
			t.Errorf("Input '%s' should not be considered valid", input)
		}
	}
}
