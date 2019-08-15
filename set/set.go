package set

import (
	"math"
	"reflect"
)

// Slice is generic map type
type Slice map[T]struct{}

// Set type with required options
type Set struct {
	data Slice
}

// CreateSet - factory method to create a new Set
// Returns pointer on created set
func CreateSet(args ...T) *Set {
	arr := make(Slice, 0)
	for arg := range iteratePack(args...) {
		arr[arg] = struct{}{}
	}
	return &Set{arr}
}

// iteratePack - range over each inbound argument with channel
func iteratePack(args ...T) Iter {
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
	return reflect.DeepEqual(set.data, s.data)
}

// Clone the set
func (set *Set) Clone() *Set {
	ret := make(Slice, 0)
	for key := range set.data {
		ret[key] = struct{}{}
	}
	return &Set{ret}
}

// iterateAndCheckPack - range and check over each inbound argument with channel
func iterateAndCheckPack(pred Predicate, args ...T) Iter {
	ch := make(Iter)

	go func() {
		for _, arg := range args {
			if pred(arg) {
				ch <- arg
			} else {
				break
			}
		}
		close(ch)
	}()

	return ch
}

// Contains returns true if all inbound arguments contains in the set
func (set *Set) Contains(args ...T) bool {
	pred := func(arg T) bool {
		_, contain := set.data[arg]
		return contain
	}

	_, contain := <-iterateAndCheckPack(pred, args...)
	return contain
}

// iterateSetsPack - range and check over each inbound set with channel
func iterateSetsPack(key T, sets []*Set) Iter {
	ch := make(Iter)

	go func() {
		for _, set := range sets {
			if set.Contains(key) {
				ch <- key
			} else {
				break
			}
		}
		close(ch)
	}()

	return ch
}

func (set *Set) Iterator() Iter {
	ch := make(Iter)

	go func() {
		for key := range set.data {
			ch <- key
		}
		close(ch)
	}()

	return ch
}

// RemoveByKeys removes items by inbound keys
func (set *Set) RemoveByKeys(args ...T) {
	for _, arg := range args {
		delete(set.data, arg)
	}
}

// Intersection returns the set which contains all the items in inbound sets,
// but no other items
func Intersection(sets ...*Set) *Set {
	// Find the shortest set
	setID := optimal{math.MaxInt64, 0}
	var idxs []int
	for idx, set := range sets {
		idxs = append(idxs, idx)
		if set.Size() < setID.length.(int) {
			setID = optimal{set.Size(), idx}
		}
	}

	// Make copy of the shortest set
	exceptedIdx := setID.idx.(int)
	ret := sets[exceptedIdx].Clone()

	// Exclude exceptedIdx
	idxs = append(idxs[:exceptedIdx], idxs[exceptedIdx+1:]...)

	// Prepare
	var setsSlice []*Set
	for _, idx := range idxs {
		setsSlice = append(setsSlice, sets[idx])
	}

	// Look for items from the shortest set in other sets
	for key := range sets[exceptedIdx].Iterator() {
		if _, contain := <-iterateSetsPack(key, setsSlice); !contain {
			ret.RemoveByKeys(key)
		}
	}

	return ret
}
