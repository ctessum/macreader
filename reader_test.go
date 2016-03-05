package macreader

import (
	"testing"
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
)

func Example() {
	// testFile is a CSV file with CR line endings.
	testFile := bytes.NewBufferString("a,b,c\r1,2,3\r").Bytes()

	// First try reading the csv file the normal way.
	// The CSV reader doesn't recognize the '\r' line ending.
	r1 := csv.NewReader(bytes.NewReader(testFile))
	lines1, err := r1.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Without macreader: %#v\n", lines1)

	// Now try reading the csv file using macreader.
	// It should work as expected.
	r2 := csv.NewReader(New(bytes.NewReader(testFile)))
	lines2, err := r2.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("With macreader: %#v\n", lines2)

	// Output: Without macreader: [][]string{[]string{"a", "b", "c\r1", "2", "3"}}
	// With macreader: [][]string{[]string{"a", "b", "c"}, []string{"1", "2", "3"}}

}


func TestCR(t *testing.T) {
	testFile := bytes.NewBufferString("a,b,c\r1,2,3\r").Bytes()

	r := csv.NewReader(New(bytes.NewReader(testFile)))
	lines, err := r.ReadAll()

	if err != nil {
		t.Errorf("An error occurred while reading the data: %v", err)
	}
	if len(lines) != 2 {
		t.Error("Wrong number of lines. Expected 2, got %d", len(lines))
	}
}

func TestLF(t *testing.T) {
	testFile := bytes.NewBufferString("a,b,c\n1,2,3\n").Bytes()

	r := csv.NewReader(New(bytes.NewReader(testFile)))
	lines, err := r.ReadAll()

	if err != nil {
		t.Errorf("An error occurred while reading the data: %v", err)
	}
	if len(lines) != 2 {
		t.Error("Wrong number of lines. Expected 2, got %d", len(lines))
	}
}

func TestCRLF(t *testing.T) {
	testFile := bytes.NewBufferString("a,b,c\r\n1,2,3\r\n").Bytes()

	r := csv.NewReader(New(bytes.NewReader(testFile)))
	lines, err := r.ReadAll()

	if err != nil {
		t.Errorf("An error occurred while reading the data: %v", err)
	}
	if len(lines) != 2 {
		t.Error("Wrong number of lines. Expected 2, got %d", len(lines))
	}
}

func TestCRInQuote(t *testing.T) {
	testFile := bytes.NewBufferString("a,\"foo,\rbar\",c\r1,\"2\r\n2\",3\r").Bytes()

	r := csv.NewReader(New(bytes.NewReader(testFile)))
	lines, err := r.ReadAll()

	if err != nil {
		t.Errorf("An error occurred while reading the data: %v", err)
	}
	if len(lines) != 2 {
		t.Error("Wrong number of lines. Expected 2, got %d", len(lines))
	}
	if strings.Contains(lines[1][1], "\n\n") {
		t.Error("The CRLF was converted to a LFLF")
	}
}