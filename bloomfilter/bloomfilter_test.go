package bloomfilter

import (
	"testing"
)

func TestNew(t *testing.T) {
	bf, err := New(32, 5)
	if err != nil && bf == nil {
		t.Errorf("Expected a BloomFilter structure, but got nil")
	}
}