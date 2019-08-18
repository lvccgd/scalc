package set

import (
	"fmt"
	"testing"
)

func TestCreateEmptySet(t *testing.T) {
	if set := CreateSet(); set == nil {
		t.Error("Create empty set fall")
	}
}

func TestCreateNonEmptySet(t *testing.T) {
	if set := CreateSet(1, 2, 3); set == nil || !set.Contains(3) {
		t.Error("Create non empty set fall")
	}
}

func TestClone(t *testing.T) {
	set := CreateSet(1, 2, 3)
	clone := set.Clone()
	if !clone.Equal(set) {
		t.Error(fmt.Errorf("Origin: %#v Clone: %#v", set, clone))
	}
}

func TestContains(t *testing.T) {
	set := CreateSet(1, 2, 3)
	if !set.Contains(1, 3) {
		t.Error(fmt.Errorf("Origin: %#v", set))
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
		set := Intersection(useCase.caseB, useCase.caseC)
		if !useCase.expect.Equal(set) {
			t.Error(fmt.Errorf("Expected: %#v Acctual: %#v", useCase.expect, set))
		}
	}
}

func TestDifference(t *testing.T) {
	var useCases = []struct {
		caseA  *Set
		caseB  *Set
		caseC  *Set
		expect *Set
	}{
		{
			caseA:  useCaseA,
			caseB:  useCaseB,
			caseC:  useCaseC,
			expect: CreateSet(1),
		},
	}

	for _, useCase := range useCases {
		set := Difference(useCase.caseA, useCase.caseB, useCase.caseC)
		if !useCase.expect.Equal(set) {
			t.Error(fmt.Errorf("Expected: %#v Acctual: %#v", useCase.expect, set))
		}
	}
}

func TestUnion(t *testing.T) {
	var useCases = []struct {
		caseA  *Set
		caseB  *Set
		expect *Set
	}{
		{
			caseA:  Difference(useCaseA, useCaseB, useCaseC),
			caseB:  Intersection(useCaseB, useCaseC),
			expect: CreateSet(1, 3, 4),
		},
		{
			caseA:  useCaseA,
			caseB:  useCaseB,
			expect: CreateSet(1, 2, 3, 4),
		},
	}

	for _, useCase := range useCases {
		set := Union(useCase.caseA, useCase.caseB)
		if !useCase.expect.Equal(set) {
			t.Error(fmt.Errorf("Expected: %#v Acctual: %#v", useCase.expect, set))
		}
	}
}
