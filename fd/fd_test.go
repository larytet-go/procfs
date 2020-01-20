package fd

import (
	"log"
	"testing"
)

func TestParsingFD(t *testing.T) {
	fds, err := New("./testfiles/fd")

	if err != nil {
		t.Fatal("Got error", err)
	}

	if fds == nil {
		t.Fatal("FD is missing")
	}
	log.Println("fds", fds)

	if len(fds) != 3 {
		t.Fatal("Expected 3 entries")
	}
}
