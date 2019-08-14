package set

import (
	"fmt"
)

// Slice is generic slice type
type Slice []T

// Set type with required options
type Set struct {
	data Slice
}

// NewSet (size int) *Set - method which create Set's DTO
// Returns pointer on created set
func newSet(size uint) *Set {
	return &Set{make(Slice, size)}
}

// CreateSet (args ...T) *Set - factory method to create a new Set
// Returns:
// - pointer on created set
// - error if it occurs
func CreateSet(args ...T) (*Set, error) {
	set := newSet(uint(len(args)))
	err := set.Copy(args...)
	return set, err
}

// GetItemByIndex returns the item of set by index
func (set *Set) GetItemByIndex(idx uint) T {
	if int(idx) >= set.Size() {
		return nil
	}
	return set.data[idx]
}

// Copy - copy additional data in set
// Returns error if it occurs
func (set *Set) Copy(args ...T) error {
	if set.Size() != len(args) {
		return fmt.Errorf("Too much arguments to copy into set. %+v", set.Size())
	}

	for idx, arg := range args {
		set.data[idx] = arg
	}

	return nil
}

// Append (args ...T) adds args to set
func (set *Set) Append(args ...T) {
	set.data = append(set.data, args...)
}

// Size returns the current set length
func (set *Set) Size() int {
	return len(set.data)
}
