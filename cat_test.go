package cat

import (
	"os"
	"testing"
)

func TestFileDoesNotExist(t *testing.T) {
	does_not_exist_file := "does_not_exist.txt"
	expected := "file not found: " + does_not_exist_file
	_, actual := Cat(does_not_exist_file)

	if actual == nil || actual.Error() != expected {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestFileUnreadable(t *testing.T) {
	undreadable_file := "test_file.txt"
	expected := "file not readable: " + undreadable_file

	// Make the file unreadable
	os.Chmod(undreadable_file, 0222)

	// Read the file
	_, actual := Cat(undreadable_file)

	// Make the file readable again
	os.Chmod(undreadable_file, 0666)

	if actual == nil || actual.Error() != expected {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestConcatenation(t *testing.T) {
	readable_file := "test_file.txt"
	expected := "aaa\n\nbbb\n\nccc"
	actual, _ := Cat(readable_file)

	if expected != actual {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}
