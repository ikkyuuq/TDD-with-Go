package dependenciesinjection

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestLogMessage(t *testing.T) {
	// bytes.Buffer which implements the io.Writer interface to capture the output.
	buffer := bytes.Buffer{}
	logMessage(&buffer, "DEBUG", "This is a debug message")

	got := buffer.String()
	want := "[DEBUG] This is a debug message"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleLogMessage_info() {
	// Print log message through console with os.Stdout
	logMessage(os.Stdout, "INFO", "This is just an info message")
	// Output: [INFO] This is just an info message
}

func ExampleLogMessage_file() {
	// Create and open a file for writing
	f, err := os.CreateTemp("", "exmaple.log")
	if err != nil {
		fmt.Println("Error can open file:", err)
		return
	}
	// Clean up after everything done
	defer os.Remove(f.Name())
	// Write log message to the file
	logMessage(f, "ERROR", "This is just a error message")
	// Read back the contents to verify
	f.Seek(0, 0) // Rewind the file
	content, _ := io.ReadAll(f)
	fmt.Print(string(content))
	// Output: [ERROR] This is just a error message
}
