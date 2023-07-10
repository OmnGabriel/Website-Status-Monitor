package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToVerifyTheReturnOfTheShowIntroFunction(t *testing.T) {
	//Arrange
	expected := "\nWebsite Monitoring Application\nIs on version: 1.0\n\n"

	//Act
	response := showIntro()

	// Assert
	if response != expected {
		t.Errorf("Response received was: %q, expected is: %q", response, expected)
	}
}

func TestToVerifyTheReturnOfTheShowMenuFunction(t *testing.T) {
	expected := "1 - Start Monitoring\n2 - Show logs\n3 - Delete logs\n0 - Leave The Program\n\n"

	response := showMenu()

	if response != expected {
		t.Errorf("Response received was: %q, expected is: %q", response, expected)
	}
}

func TestToVerifyTheReturnOfTheReadCommandFunctionIsAInt(t *testing.T) {
	var expected int

	response := readCommand()

	if response != expected {
		t.Errorf("Response received was: %d, expected is: %d", response, expected)
	}
}

func TestToVerifyTheLoopForMonitoringFunction(t *testing.T) {
	expected := 0

	response := loopForMonitoring(0)

	if response != expected {
		t.Errorf("Response received was: %d, expected is: %d", response, expected)
	}
}

func TestToVerifyReadingFilesSitesFunction(t *testing.T) {
	var expected []string

	response := readingFilesSites()

	fmt.Println()
	if cmp.Equal(response, expected) == true {
		t.Errorf("Response received was: %q, expected is: %q", response, expected)
	}
}

func TestToVerifyErrorHandlingFunction(t *testing.T) {
	expected := "\x1d"
	var err error

	response := errorHandling(err)
	if response != expected {
		t.Errorf("Response received was: %q, expected is: %q", response, expected)

	}

}

func TestToVerifyStartMonitoringLoopFunction(t *testing.T) {
	expected := 0

	response := startMonitoring()

	if response != expected {
		t.Errorf("Response received was: %d, expected is: %d", response, expected)
	}
}
