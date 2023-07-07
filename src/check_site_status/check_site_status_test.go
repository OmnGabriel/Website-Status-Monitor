package main

import (
	"testing"
)

func TestToVerifyTheReturnOfTheShowIntroFunction(t *testing.T) {
	//Arrange
	response := "\nWebsite Monitoring Application\nIs on version: 1.0\n\n"

	//Act
	espected := showIntro()

	// Assert
	if response != espected {
		t.Errorf("Response received was: %q, expected is: %q", response, espected)
	}
}

func TestToVerifyTheReturnOfTheShowMenuFunction(t *testing.T) {
	response := "1 - Start Monitoring\n2 - Show logs\n3 - Delete logs\n0 - Leave The Program\n\n"

	espected := showMenu()

	if response != espected {
		t.Errorf("Response received was: %q, expected is: %q", response, espected)
	}
}

func TestToVerifyTheReturnOfTheReadCommandFunctionIsAInt(t *testing.T) {
	var response int

	espected := readCommand()

	if response != espected {
		t.Errorf("Response received was: %d, expected is: %d", response, espected)
	}
}

func TestToVerifyTheLoopForMonitoringFunction(t *testing.T) {
	response := 0

	espected := loopForMonitoring(0)

	if response != espected {
		t.Errorf("Response received was: %d, expected is: %d", response, espected)
	}
}
