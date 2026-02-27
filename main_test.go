//
// package main
//
// For the test file, either the package name can be `main` (same as the main program) or 
// `main_test` (a separate test package).
//
// Normally, you cannot have two packages in the same directory. However, for testing 
// purposes, Go allows you to have a separate test package (e.g., `main_test`) in the same 
// directory as the main package (`main`). This is a common practice to keep test code 
// separate from production code. In this case, we will use `main_test` for our test file 
// to follow best practices.
package main_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	main "github.com/hgbbus/ChineseZodiacGo"
)

func TestMain(t *testing.T) {
	t.Log("Starting main function test...")

	// ---- Save original stdin and stdout ----
	oldStdin := os.Stdin
	oldStdout := os.Stdout

	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	// ---- Mock stdin and send input ----
	inReader, inWriter, _ := os.Pipe()
	os.Stdin = inReader
	//
	input := "1\n1984\n2\n2020\n99\n0\n"
	inWriter.Write([]byte(input))
	inWriter.Close()

	// ---- Mock stdout to capture output ----
	outReader, outWriter, _ := os.Pipe()
	os.Stdout = outWriter

	// ---- Run program ----
	main.Main()
	outWriter.Close()	// close the writer to signal end of output

	// ---- Read captured output ----
	var output bytes.Buffer
	io.Copy(&output, outReader)
	t.Log("Captured output:\n", output.String())

	// ---- Basic assertion to check if the output contains expected strings ----
	if !strings.Contains(output.String(), "Welcome") {
		t.Error("Expected welcome message not found in output")
	}
	if !strings.Contains(output.String(), "Please select an option:") {
		t.Error("Expected menu prompt not found in output")
	}
	if !strings.Contains(output.String(), "Enter your choice ") {
		t.Error("Expected choice prompt not found in output")
	}
	if !strings.Contains(output.String(), "Invalid choice") {
		t.Error("Expected invalid choice message not found in output")
	}
	if !strings.Contains(output.String(), "Exiting the program.") {
		t.Error("Expected exit message not found in output")
	}

	t.Log("Main function executed successfully.")
}
