package files

import (
	"testing"
)

func TestListFiles(t *testing.T) {
	entries, err := ListFiles()
	if err != nil {
		t.Errorf("ListFiles() error = %v", err)
	}
	for _, entry := range entries {
		t.Logf("Found entry: %s", entry.Name())
	}
}
