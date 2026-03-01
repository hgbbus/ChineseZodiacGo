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
	"fmt"
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
	//t.Log("Captured output:\n", output.String())

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

func TestGetElement(t *testing.T) {
	tests := []struct {
		stemIndex int
		expected  string
	}{
		{0, "Wood"},
		{1, "Wood"},
		{2, "Fire"},
		{3, "Fire"},
		{4, "Earth"},
		{5, "Earth"},
		{6, "Metal"},
		{7, "Metal"},
		{8, "Water"},
		{9, "Water"},
	}

	for _, test := range tests {
		// Use t.Run to create subtests for better reporting
		t.Run(fmt.Sprintf("stemIndex=%d", test.stemIndex), func(t *testing.T) {
			result := main.GetElement(test.stemIndex)
			if result != test.expected {
				t.Errorf("GetElement(%d) = %s; expected %s", test.stemIndex, result, test.expected)
			}
		})
	}
}

func TestGetYinYang(t *testing.T) {
	tests := []struct {
		stemIndex int
		expected  string
	}{
		{0, "Yang"},
		{1, "Yin"},
		{2, "Yang"},
		{3, "Yin"},
		{4, "Yang"},
		{5, "Yin"},
		{6, "Yang"},
		{7, "Yin"},
		{8, "Yang"},
		{9, "Yin"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("stemIndex=%d", test.stemIndex), func(t *testing.T) {
			result := main.GetYinYang(test.stemIndex)
			if result != test.expected {
				t.Errorf("GetYinYang(%d) = %s; expected %s", test.stemIndex, result, test.expected)
			}
		})
	}
}

func TestGetZodiacInfo(t *testing.T) {
	tests := []struct {
		year     int
		expected [5]string
	}{
		{1984, [5]string{"Jia", "Zi", "Wood", "Yang", "Rat"}},
		{1985, [5]string{"Yi", "Chou", "Wood", "Yin", "Ox"}},
		{1986, [5]string{"Bing", "Yin", "Fire", "Yang", "Tiger"}},
		{1987, [5]string{"Ding", "Mao", "Fire", "Yin", "Rabbit"}},
		{1988, [5]string{"Wu", "Chen", "Earth", "Yang", "Dragon"}},
		{1989, [5]string{"Ji", "Si", "Earth", "Yin", "Snake"}},
		{1990, [5]string{"Geng", "Wu", "Metal", "Yang", "Horse"}},
		{1991, [5]string{"Xin", "Wei", "Metal", "Yin", "Goat"}},
		{1992, [5]string{"Ren", "Shen", "Water", "Yang", "Monkey"}},
		{1993, [5]string{"Gui", "You", "Water", "Yin", "Rooster"}},
		{1994, [5]string{"Jia", "Xu", "Wood", "Yang", "Dog"}},
		{1995, [5]string{"Yi", "Hai", "Wood", "Yin", "Pig"}},
		{1996, [5]string{"Bing", "Zi", "Fire", "Yang", "Rat"}},
		{1997, [5]string{"Ding", "Chou", "Fire", "Yin", "Ox"}},
		{1998, [5]string{"Wu", "Yin", "Earth", "Yang", "Tiger"}},
		{1999, [5]string{"Ji", "Mao", "Earth", "Yin", "Rabbit"}},
		{2000, [5]string{"Geng", "Chen", "Metal", "Yang", "Dragon"}},
		{2001, [5]string{"Xin", "Si", "Metal", "Yin", "Snake"}},
		{2002, [5]string{"Ren", "Wu", "Water", "Yang", "Horse"}},
		{2003, [5]string{"Gui", "Wei", "Water", "Yin", "Goat"}},
		{2004, [5]string{"Jia", "Shen", "Wood", "Yang", "Monkey"}},
		{2005, [5]string{"Yi", "You", "Wood", "Yin", "Rooster"}},
		{2006, [5]string{"Bing", "Xu", "Fire", "Yang", "Dog"}},
		{2007, [5]string{"Ding", "Hai", "Fire", "Yin", "Pig"}},
		{2008, [5]string{"Wu", "Zi", "Earth", "Yang", "Rat"}},
		{2009, [5]string{"Ji", "Chou", "Earth", "Yin", "Ox"}},
		{2010, [5]string{"Geng", "Yin", "Metal", "Yang", "Tiger"}},
		{2011, [5]string{"Xin", "Mao", "Metal", "Yin", "Rabbit"}},
		{2012, [5]string{"Ren", "Chen", "Water", "Yang", "Dragon"}},
		{2013, [5]string{"Gui", "Si", "Water", "Yin", "Snake"}},
		{2014, [5]string{"Jia", "Wu", "Wood", "Yang", "Horse"}},
		{2015, [5]string{"Yi", "Wei", "Wood", "Yin", "Goat"}},
		{2016, [5]string{"Bing", "Shen", "Fire", "Yang", "Monkey"}},
		{2017, [5]string{"Ding", "You", "Fire", "Yin", "Rooster"}},
		{2018, [5]string{"Wu", "Xu", "Earth", "Yang", "Dog"}},
		{2019, [5]string{"Ji", "Hai", "Earth", "Yin", "Pig"}},
		{2020, [5]string{"Geng", "Zi", "Metal", "Yang", "Rat"}},
		{2021, [5]string{"Xin", "Chou", "Metal", "Yin", "Ox"}},
		{2022, [5]string{"Ren", "Yin", "Water", "Yang", "Tiger"}},
		{2023, [5]string{"Gui", "Mao", "Water", "Yin", "Rabbit"}},
		{2024, [5]string{"Jia", "Chen", "Wood", "Yang", "Dragon"}},
		{2025, [5]string{"Yi", "Si", "Wood", "Yin", "Snake"}},
		{2026, [5]string{"Bing", "Wu", "Fire", "Yang", "Horse"}},
		{2027, [5]string{"Ding", "Wei", "Fire", "Yin", "Goat"}},
		{2028, [5]string{"Wu", "Shen", "Earth", "Yang", "Monkey"}},
		{2029, [5]string{"Ji", "You", "Earth", "Yin", "Rooster"}},
		{2030, [5]string{"Geng", "Xu", "Metal", "Yang", "Dog"}},
		{2031, [5]string{"Xin", "Hai", "Metal", "Yin", "Pig"}},
		{2032, [5]string{"Ren", "Zi", "Water", "Yang", "Rat"}},
		{2033, [5]string{"Gui", "Chou", "Water", "Yin", "Ox"}},
		{2034, [5]string{"Jia", "Yin", "Wood", "Yang", "Tiger"}},
		{2035, [5]string{"Yi", "Mao", "Wood", "Yin", "Rabbit"}},
		{2036, [5]string{"Bing", "Chen", "Fire", "Yang", "Dragon"}},
		{2037, [5]string{"Ding", "Si", "Fire", "Yin", "Snake"}},
		{2038, [5]string{"Wu", "Wu", "Earth", "Yang", "Horse"}},
		{2039, [5]string{"Ji", "Wei", "Earth", "Yin", "Goat"}},
		{2040, [5]string{"Geng", "Shen", "Metal", "Yang", "Monkey"}},
		{2041, [5]string{"Xin", "You", "Metal", "Yin", "Rooster"}},
		{2042, [5]string{"Ren", "Xu", "Water", "Yang", "Dog"}},
		{2043, [5]string{"Gui", "Hai", "Water", "Yin", "Pig"}},
		{2044, [5]string{"Jia", "Zi", "Wood", "Yang", "Rat"}},
	}
	
	for _, test := range tests {
		t.Run(fmt.Sprintf("year=%d", test.year), func(t *testing.T) {
			stem, branch, element, yinYang, animal := main.GetZodiacInfo(test.year)
			result := [5]string{stem, branch, element, yinYang, animal}
			if result != test.expected {
				t.Errorf("GetZodiacInfo(%d) = %v; expected %v", test.year, result, test.expected)
			}
		})
	}
}