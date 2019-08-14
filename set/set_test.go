package set

import (
	"testing"
)

func TestGetItemByIndex(t *testing.T) {
	set := new(Set)
	if set.GetItemByIndex(1) != nil {
		t.Error("Out of range missed")
	}
}

func TestCopyNewItemsIntoSet(t *testing.T) {
	set := new(Set)
	set.Copy(1, 2, 3)
	if err := set.Copy(1, 2, 3); err != nil || set.GetItemByIndex(2) != 3 {
		t.Error(err)
	}
}

func TestCreateEmptySet(t *testing.T) {
	if set, err := CreateSet(); set == nil {
		t.Error(err)
	}
}

func TestCreateNonEmptySet(t *testing.T) {
	if _, err := CreateSet(1, 2, 3); err != nil {
		t.Error(err)
	}
}
