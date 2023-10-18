package bloomfilter

import (
	"testing"
)

func TestBloomFilter_InsertAndContains(t *testing.T) {
	bf, err := NewBloomFilter(128, 3)
	if err != nil {
		t.Errorf("Failed to create Bloom Filter: %v", err)
	}

	items := []string{"apple", "banana", "cherry", "date"}

	for _, item := range items {
		err := bf.Insert(item)
		if err != nil {
			t.Errorf("Insertion failed: %v", err)
		}

		exists, err := bf.Contains(item)
		if err != nil {
			t.Errorf("Contains failed: %v", err)
		}

		if !exists {
			t.Errorf("Item %s should exist in the Bloom Filter", item)
		}
	}
}

func TestBloomFilter_FalsePositiveRate(t *testing.T) {
	bf, err := NewBloomFilter(128, 3)
	if err != nil {
		t.Errorf("Failed to create Bloom Filter: %v", err)
	}

	// Insert items
	items := []string{"apple", "banana", "cherry", "date"}
	for _, item := range items {
		_ = bf.Insert(item)
	}

	// Check false positive rate
	fpr, err := bf.FalsePositiveRate()
	if err != nil {
		t.Errorf("FalsePositiveRate failed: %v", err)
	}

	if fpr < 0 || fpr > 1 {
		t.Errorf("Invalid false positive rate: %f", fpr)
	}
}

func TestBloomFilter_Union(t *testing.T) {
	bf1, err := NewBloomFilter(128, 3)
	if err != nil {
		t.Errorf("Failed to create Bloom Filter: %v", err)
	}

	bf2, err := NewBloomFilter(128, 3)
	if err != nil {
		t.Errorf("Failed to create Bloom Filter: %v", err)
	}

	items1 := []string{"apple", "banana", "cherry"}
	items2 := []string{"cherry", "date", "fig"}

	for _, item := range items1 {
		_ = bf1.Insert(item)
	}

	for _, item := range items2 {
		_ = bf2.Insert(item)
	}

	union, err := bf1.Union(bf2)
	if err != nil {
		t.Errorf("Union failed: %v", err)
	}

	for _, item := range items1 {
		exists, err := union.Contains(item)
		if err != nil {
			t.Errorf("Contains failed: %v", err)
		}

		if !exists {
			t.Errorf("Item %s should exist in the union Bloom Filter", item)
		}
	}

	for _, item := range items2 {
		exists, err := union.Contains(item)
		if err != nil {
			t.Errorf("Contains failed: %v", err)
		}

		if !exists {
			t.Errorf("Item %s should exist in the union Bloom Filter", item)
		}
	}
}

func TestBloomFilter_Intersection(t *testing.T) {
	bf1, err := NewBloomFilter(128, 3)
	if err != nil {
		t.Errorf("Failed to create Bloom Filter: %v", err)
	}

	bf2, err := NewBloomFilter(128, 3)
	if err != nil {
		t.Errorf("Failed to create Bloom Filter: %v", err)
	}

	items1 := []string{"apple", "banana", "cherry"}
	items2 := []string{"cherry", "date", "fig"}

	for _, item := range items1 {
		_ = bf1.Insert(item)
	}

	for _, item := range items2 {
		_ = bf2.Insert(item)
	}

	intersection, err := bf1.Intersection(bf2)
	if err != nil {
		t.Errorf("Intersection failed: %v", err)
	}

	for _, item := range items1 {
		exists, err := intersection.Contains(item)
		if err != nil {
			t.Errorf("Contains failed: %v", err)
		}

		if exists && item != "cherry" {
			t.Errorf("Item %s should not exist in the intersection Bloom Filter", item)
		}
	}

	for _, item := range items2 {
		exists, err := intersection.Contains(item)
		if err != nil {
			t.Errorf("Contains failed: %v", err)
		}

		if exists && item != "cherry" {
			t.Errorf("Item %s should not exist in the intersection Bloom Filter", item)
		}
	}
}

func TestBloomFilter_NumberOfItems(t *testing.T) {
	bf, err := NewBloomFilter(128, 3)
	if err != nil {
		t.Errorf("Failed to create Bloom Filter: %v", err)
	}

	items := []string{"apple", "banana", "cherry", "date", "fig"}

	for _, item := range items {
		_ = bf.Insert(item)
	}

	count, err := bf.NumberOfItems()
	if err != nil {
		t.Errorf("NumberOfItems failed: %v", err)
	}

	if count != len(items) {
		t.Errorf("Incorrect number of items. Expected: %d, Got: %d", len(items), count)
	}
}
