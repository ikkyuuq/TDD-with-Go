package dependenciesinjection

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLogMessage(t *testing.T) {
	t.Run("capture the output with bytes.Buffer", func(t *testing.T) {
		// bytes.Buffer which implements the io.Writer interface to capture the output.
		buffer := bytes.Buffer{}
		logMessage(&buffer, "DEBUG", "This is a debug message")

		got := buffer.String()
		want := "[DEBUG] This is a debug message"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("the internet", func(t *testing.T) {
		// Create new http server with httptest that take hanlderfunc with LogMessageHandler
		server := httptest.NewServer(http.HandlerFunc(LogMessageHandler))
		// Close server
		defer server.Close()
		// Mockup url for testing with server.URL following with query
		url := fmt.Sprintf("%s?level=INFO&message=This+a+INFO+message+from+internet", server.URL)

		// GET on specific URL
		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("Failed to make a request: %v", err)
		}
		// Close resp after reading from it
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %v", resp.StatusCode)
		}

		// Create a recorder that take body field
		body := httptest.NewRecorder().Body
		// Read the response body
		body.ReadFrom(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		// Get a string from *bytes.Buffer (body) that read body of resp
		got := body.String()
		want := "[INFO] This a INFO message from internet"

		if got != want {
			t.Errorf("got %q want %q", server.URL, want)
		}
	})
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

func ExmapleLogMessage_internet() {
	server := httptest.NewServer(http.HandlerFunc(LogMessageHandler))
	defer server.Close()

	url := fmt.Sprintf("%s?level=INFO&message=This+a+INFO+message+from+internet", server.URL)
	resp, _ := http.Get(url)
	resp.Body.Close()

	body := httptest.NewRecorder().Body
	body.Reset()
	body.ReadFrom(resp.Body)

	fmt.Println(body.String())
	// Output: [INFO] This a INFO message from internet
}
