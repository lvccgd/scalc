package set

import (
	"math"
	"reflect"
)

// Slice is generic slice type
type Slice []T

// Set type with required options
type Set struct {
	data Slice
}

// CreateSet - factory method to create a new Set
// Returns pointer on created set
func CreateSet(args ...T) *Set {
	arr := make(Slice, 0, 1)
	for item := range newSet(args...) {
		arr = append(arr, item)
	}
	return &Set{arr}
}

// newSet - range over each inbound argument with channel
func newSet(args ...T) Iter {
	ch := make(Iter)
	go func() {
		for _, arg := range args {
			ch <- arg
		}
		close(ch)
	}()

	return ch
}

// GetItemByIndex returns the item of set by index
func (set *Set) GetItemByIndex(idx uint) T {
	if int(idx) >= set.Size() {
		return nil
	}
	return set.data[idx]
}

// Size returns the current set length
func (set *Set) Size() int {
	return len(set.data)
}

// Equal returns true if
func (set *Set) Equal(s *Set) bool {
	return reflect.DeepEqual(set.data, s)
}

// Intersection returns the set which contains all the elements of all items
// from inbound sets, but no other elements
func Intersection(sets ...*Set) (*Set, error) {
	// Find the shortest set
	setID := optimal{math.MaxInt64, 0}
	for idx, set := range sets {
		if set.Size() < setID.length.(int) {
			setID = optimal{set.Size(), idx}
		}
	}

	// Look for elements from the shortest set in other sets

	return nil, nil
}
