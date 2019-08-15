package set

import (
	"fmt"
	"testing"
)

func TestGetItemByIndex(t *testing.T) {
	set := CreateSet()
	if set.GetItemByIndex(0) != nil {
		t.Error("Out of range missed")
	}
}

func TestCopyNewItemsIntoSet(t *testing.T) {
	if set := CreateSet(1, 2, 3); set == nil || set.GetItemByIndex(2) != 3 {
		t.Error("Create non empty set fall")
	}
}

func TestCreateEmptySet(t *testing.T) {
	if set := CreateSet(); set == nil {
		t.Error("Create empty set fall")
	}
}

func TestCreateNonEmptySet(t *testing.T) {
	if set := CreateSet(1, 2, 3); set == nil || set.GetItemByIndex(2) != 3 {
		t.Error("Create non empty set fall")
	}
}

var useCaseA = CreateSet(1, 2, 3)
var useCaseB = CreateSet(2, 3, 4)
var useCaseC = CreateSet(3, 4, 5)

func TestIntersection(t *testing.T) {
	var useCases = []struct {
		caseB  *Set
		caseC  *Set
		expect *Set
	}{
		{
			caseB:  useCaseB,
			caseC:  useCaseC,
			expect: CreateSet(3, 4),
		},
	}

	for _, useCase := range useCases {
		set, err := Intersection(useCase.caseB, useCase.caseC)
		if err != nil {
			t.Error(err)
		}
		if !useCase.expect.Equal(set) {
			t.Error(fmt.Errorf("Expected: %#v Acctual: %#v", useCase.expect, set))
		}
	}
}
